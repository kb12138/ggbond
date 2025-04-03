package main

import (
	"fmt"
	"time"
)

// 定义一个结构体
type Rectangle struct {
	width  float64
	height float64
}

// 定义一个使用指针接收者的方法
func (r *Rectangle) SetWidth(width float64) {
	r.width = width
}

var xx = 1

func main() {
	rect := Rectangle{width: 10, height: 5}
	rect.SetWidth(20)
	fmt.Printf("修改后的矩形宽度: %.2f\n", rect.width)

	fmt.Println(fmt.Sprintf("dd__%s__%d", "eeee", 5654))

	fmt.Println(xx)
	xx = 4
	fmt.Println(xx)

	go sayHello()
	sayWorld()

	time.Sleep(time.Second)

	var intArray [10]int
	intArray[0] = 1
	fmt.Println(intArray)

	floatArray := [11]float32{}
	floatArray[0] = 1.03
	fmt.Println(floatArray)

	var ptr *int

	ptr = &intArray[0]

	fmt.Println(ptr)

	var s = make([]int, 0)
	_ = s

}

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(time.Millisecond * 100)
	}
}

func sayWorld() {
	for i := 0; i < 5; i++ {
		fmt.Println("World")
		time.Sleep(time.Millisecond * 100)
	}
}
