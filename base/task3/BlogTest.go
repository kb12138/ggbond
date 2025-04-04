package task3

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1. 用户模型（基础字段 + 一对多关联）
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`                         // 主键
	CreatedAt time.Time `json:"created_at"`                                   // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                                   // 更新时间
	Username  string    `gorm:"uniqueIndex;not null;size:50" json:"username"` // 唯一用户名
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`            // 唯一邮箱
	Password  string    `gorm:"not null" json:"-"`                            // 密码哈希（敏感字段）
	// 一对多：用户发布的文章（仅保留外键，移除级联）
	Posts      []Post `gorm:"foreignKey:AuthorID" json:"posts,omitempty"` // 外键指向 AuthorID
	PostsCount uint   `gorm:"default:0;not null" json:"posts_count"`      // 文章数量统计

}

// 2. 文章模型（双向关联，无自动级联）
type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`           // 主键
	CreatedAt time.Time `json:"created_at"`                     // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                     // 更新时间
	Title     string    `gorm:"not null;size:100" json:"title"` // 文章标题
	Content   string    `gorm:"not null" json:"content"`        // 文章内容
	// 多对一：文章作者（外键关联 User.ID）
	AuthorID uint `gorm:"not null" json:"author_id"`             // 外键字段
	Author   User `gorm:"references:ID" json:"author,omitempty"` // 关联用户
	// 一对多：文章的评论（仅保留外键，移除级联）
	Comments      []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"` // 外键指向 PostID
	CommentStatus string    `gorm:"default:'有评论';size:20" json:"comment_status"` // 评论状态

}

// 3. 评论模型（多对多关联）
type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`             // 主键
	CreatedAt time.Time `json:"created_at"`                       // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                       // 更新时间
	Content   string    `gorm:"not null;size:500" json:"content"` // 评论内容
	// 多对一：评论所属文章（外键关联 Post.ID）
	PostID uint `gorm:"not null" json:"post_id"`             // 外键字段
	Post   Post `gorm:"references:ID" json:"post,omitempty"` // 关联文章
	// 多对一：评论作者（外键关联 User.ID）
	UserID uint `gorm:"not null" json:"user_id"`             // 外键字段
	User   User `gorm:"references:ID" json:"user,omitempty"` // 关联用户
}

// 注册模型钩子（在初始化DB后调用）
func setupHooks(db *gorm.DB) {
	db.Callback().Create().After("gorm:create").Register("post:after_create", updateUserPostsCount)
	db.Callback().Delete().After("gorm:delete").Register("comment:after_delete", checkPostCommentStatus)
}

// 钩子函数：Post创建后，原子性增加用户文章数
func updateUserPostsCount(tx *gorm.DB) {
	// 1. 从钩子上下文获取新创建的Post
	var post Post
	if err := tx.First(&post).Error; err != nil {
		return // 无主键，跳过
	}

	// 2. 原子性更新用户文章数（防止并发竞争）
	if err := tx.Model(&User{}).
		Where("id = ?", post.AuthorID).
		UpdateColumn("posts_count", gorm.Expr("posts_count + ?", 1)).
		Error; err != nil {
		tx.AddError(err) // 回滚事务
	}
}

// 钩子函数：Comment删除后，检查文章评论数
func checkPostCommentStatus(tx *gorm.DB) {
	// 1. 从钩子上下文获取被删除的Comment
	var comment Comment
	if err := tx.First(&comment).Error; err != nil {
		return
	}

	// 2. 统计文章剩余评论数（排除软删除）
	var commentCount int64
	tx.Model(&Comment{}).
		Where("post_id = ? AND deleted_at IS NULL", comment.PostID).
		Count(&commentCount)

	// 3. 若评论数为0，更新文章状态（使用UpdateColumn避免全量更新）
	if commentCount == 0 {
		if err := tx.Model(&Post{}).
			Where("id = ?", comment.PostID).
			UpdateColumn("comment_status", "无评论").
			Error; err != nil {
			tx.AddError(err) // 回滚事务
		}
	}
}

func initDB() {
	// 4. 数据库连接与迁移
	dsn := "root:password@tcp(127.0.0.1:3306)/blog_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	// 5. 自动迁移（创建表+外键+索引）
	if err := db.AutoMigrate(
		&User{},
		&Post{},
		&Comment{},
	); err != nil {
		panic("表迁移失败: " + err.Error())
	}

	// 2. 注册钩子
	setupHooks(db)

	// 6. 打印表结构（调试用）
	sqlDB, _ := db.DB()
	tables, _ := sqlDB.Query("SHOW TABLES")
	fmt.Println("成功创建以下表：")
	for tables.Next() {
		var table string
		tables.Scan(&table)
		fmt.Println("-", table)
	}
}

func main1() {
	dsn := "root:password@tcp(127.0.0.1:3306)/blog_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	// 1. 查询某个用户发布的所有文章及其对应的评论信息（示例用户ID=1）
	queryUserPostsAndComments(db, 1)

	// 2. 查询评论数量最多的文章信息
	queryMostCommentedPost(db)
}

// 查询用户发布的文章及评论
func queryUserPostsAndComments(db *gorm.DB, userID uint) {
	var user User
	// 预加载文章及文章的评论
	if err := db.Preload("Posts.Comments").First(&user, userID).Error; err != nil {
		fmt.Println("查询用户失败:", err)
		return
	}
	fmt.Printf("用户ID %d 的文章数量: %d\n", user.ID, len(user.Posts))
	for _, post := range user.Posts {
		fmt.Printf("  文章ID: %d, 标题: %s, 评论数量: %d\n",
			post.ID, post.Title, len(post.Comments))
	}
}

// 查询评论数量最多的文章
func queryMostCommentedPost(db *gorm.DB) {
	var post Post
	// 通过子查询统计评论数并排序
	if err := db.Preload("Comments").
		Order("(SELECT COUNT(*) FROM comments WHERE comments.post_id = posts.id) DESC").
		Limit(1).
		First(&post).Error; err != nil {
		fmt.Println("查询文章失败:", err)
		return
	}
	fmt.Printf("\n评论最多的文章ID: %d, 标题: %s, 评论数: %d\n",
		post.ID, post.Title, len(post.Comments))
}
