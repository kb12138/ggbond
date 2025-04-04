package task3

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {

}

// 定义与表结构匹配的Employee结构体
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func run() {
	// 1. 连接数据库（替换为实际DSN）
	db, err := sqlx.Connect("mysql", "user:pass@tcp(127.0.0.1:3306)/company?parseTime=true")
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close() // 确保连接关闭

	// 2. 查询技术部员工（带预处理和切片映射）
	query := `
        SELECT id, name, department, salary 
        FROM employees 
        WHERE department = ? 
        ORDER BY salary DESC
    `
	techEmployees := []Employee{}
	if err := db.Select(&techEmployees, query, "技术部"); err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	// 3. 输出结果（带格式化）
	fmt.Println("技术部员工列表（共", len(techEmployees), "人）:")
	for _, emp := range techEmployees {
		fmt.Printf("  %d. %-8s 部门: %-6s 工资: %.2f\n",
			emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	// 2. 查询工资最高的员工（单个结构体）
	highestEmp, err := getHighestPaidEmployee(db)
	if err != nil {
		log.Println("最高工资查询:", err)
	} else {
		fmt.Printf("\n最高工资员工: %s（%.2f元）\n", highestEmp.Name, highestEmp.Salary)
	}

}

// 2. 查询工资最高的员工（单个结构体）
func getHighestPaidEmployee(db *sqlx.DB) (*Employee, error) {
	query := `
        SELECT id, name, department, salary 
        FROM employees 
        ORDER BY salary DESC 
        LIMIT 1
    `
	var emp Employee
	err := db.Get(&emp, query)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("没有员工记录")
	}
	if err != nil {
		return nil, fmt.Errorf("查询最高工资失败: %v", err)
	}
	return &emp, nil
}
