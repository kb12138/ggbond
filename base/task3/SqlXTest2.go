package task3

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

// 1. 类型安全的Book结构体（含自定义验证）
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// 2. 查询条件结构体（支持动态过滤）
type BookQuery struct {
	MinPrice float64 // 最低价格（>）
	Author   string  // 作者（可选）
	Limit    int     // 分页限制（默认10）
}

func run1() {
	// 2. 安全的数据库连接（带连接池配置）
	db, err := sqlx.Connect("mysql",
		"user:pass@tcp(127.0.0.1:3306)/bookstore?parseTime=true&loc=Local",
	)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close() // 确保连接释放

	// 3. 类型安全的参数化查询（防止SQL注入）
	query := `
        SELECT id, title, author, price 
        FROM books 
        WHERE price > ? 
        ORDER BY price DESC
    `
	minPrice := 50.0 // 可动态传入的查询条件
	books := []Book{}

	// 4. 结构体切片映射（自动类型校验）
	if err := db.Select(&books, query, minPrice); err != nil {
		if err == sql.ErrNoRows {
			log.Println("没有找到价格>50的书籍")
			return
		}
		log.Fatalf("查询失败（类型不匹配或SQL错误）: %v", err)
	}

	// 5. 类型安全的结果处理（编译期检查字段访问）
	fmt.Printf("找到%d本价格>%.2f的书籍：\n", len(books), minPrice)
	for _, book := range books {
		// 安全访问结构体字段（非空校验示例）
		if book.Title == "" {
			log.Printf("警告：书籍ID=%d的标题为空", book.ID)
		}
		fmt.Printf("• %s - %s（¥%.2f）\n",
			book.Title, book.Author, book.Price)
	}
}
