package main

import "fmt"

//
//func main() {
//	var input *int
//	var i = 10
//	input = &i
//	fmt.Println(*input)
//
//	fmt.Println(add10(input))
//	fmt.Println(*input)
//
//	var s []int = []int{1, 2, 3, 4}
//	var sp *[]int = &s
//	fmt.Println(*sp)
//	multiply2(sp)
//	fmt.Println(*sp)
//
//}

func add10(input *int) int {
	*input += 10
	return *input
}

func multiply2(input *[]int) []int {
	var s = *input
	for i, num := range s {
		s[i] = num * 2
	}
	fmt.Println("multiply2: %s", *input)
	return *input
}
