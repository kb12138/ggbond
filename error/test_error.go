package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("this is a error")
	fmt.Println("hello world")

	fmt.Println(err)
}
