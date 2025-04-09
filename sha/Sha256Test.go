package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	doSha256("sha256 this string")
	doSha256("sha256 this string0")
	doSha256("sha256 this string0")
}

func doSha256(s string) {
	// 我们在这里生成一个新的哈希
	h := sha256.New()
	// Write方法期望字节数据。如果您有一个字符串s，可以使用[]byte(s)将其强制转换为字节。
	h.Write([]byte(s))
	// 这将最终的哈希结果作为字节切片获取。Sum的参数可用于追加到现有的字节切片：通常不需要这样做。
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
