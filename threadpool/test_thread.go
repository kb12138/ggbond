package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//thread1();
	thread2()

}

func thread2() {
	waitGroup := sync.WaitGroup{}
	go cal(&waitGroup)
	waitGroup.Add(1)
	go cal(&waitGroup)
	waitGroup.Add(1)
	waitGroup.Wait()

	fmt.Println("ALL worker HAS done")
}

func cal(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("cal.......")
	time.Sleep(time.Second)
}

func thread1() {
	var ch = make(chan int, 2)

	ch <- 1
	ch <- 2

	int1 := <-ch
	var int2 = <-ch

	fmt.Println(int1)
	fmt.Println(int2)
}
