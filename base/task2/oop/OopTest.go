package main

import "fmt"

func main() {

}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (r Rectangle) Area() {

}
func (r Rectangle) Perimeter() {

}

type Circle struct {
}

func (r Circle) Area() {

}
func (r Circle) Perimeter() {

}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) printInfo() {
	fmt.Printf("员工信息：\n")
	fmt.Printf("  姓名: %s\n", e.Name) // 直接访问嵌入的 Name 字段
	fmt.Printf("  年龄: %d\n", e.Age)  // 直接访问嵌入的 Age 字段
	fmt.Printf("  工号: %s\n", e.EmployeeID)
}
