package main

import (
	"fmt"
	"time"
)

func main() {
	runWithBuff()
}

func run() {
	var ch = make(chan int)
	go func() {
		for i := range 10 {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for i := range ch {
			fmt.Printf("任务: %6v\n", i)
		}
	}()

	time.Sleep(time.Second)

}

func runWithBuff() {
	var ch = make(chan int, 10)
	go func() {
		for i := range 100 {
			ch <- i
			fmt.Printf("生产者发送: %3d （通道剩余容量: %d）\n", i, cap(ch)-len(ch))
			time.Sleep(time.Millisecond * 100)
		}
		close(ch)
	}()

	go func() {
		for i := range ch {
			fmt.Printf("消费者 接收: %3d （通道剩余元素: %d）\n", i, len(ch))
			time.Sleep(time.Millisecond * 10)
		}
	}()

	time.Sleep(10 * time.Second)

}

func producer(chan int) {

}

func consumer() {

}
