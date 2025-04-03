package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	run()
	runWithAtomic()
}

func run() {
	var lock = sync.Mutex{}
	var counter = 0
	var wg sync.WaitGroup
	var start = time.Now()

	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i2 := range 1000 {
				lock.Lock()
				counter++
				fmt.Printf("线程%d 接收:第 %d 次加锁 （通道剩余元素: %d）\n", i, i2)
				lock.Unlock()
			}
		}()
	}
	wg.Wait()

	fmt.Printf("执行完成 counter （%d）, time cost: %d \n", counter, time.Since(start).Milliseconds())

}

func runWithAtomic() {
	var counter = int64(0)
	var wg sync.WaitGroup
	var start = time.Now()

	for i := range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i2 := range 1000 {
				atomic.AddInt64(&counter, 1)
				fmt.Printf("线程%d 接收:第 %d 次加锁 （通道剩余元素: %d）\n", i, i2)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("执行完成 counter （%d）, time cost: %d \n", counter, time.Since(start).Milliseconds())

}
