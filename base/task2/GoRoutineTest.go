package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	go printBy(10, 1)
	go printBy(10, 0)

	time.Sleep(time.Second)

	taskScheduler()

}

func printBy(max int, step int) {
	for i := 1; i <= max; i++ {
		if i%2 == step {
			fmt.Printf("routineID:[%d],  ==> %d\n", runtime.NumGoroutine(), i)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// 任务调度器
type TaskStat struct {
	Name      string
	Duration  time.Duration
	Completed bool
}

func taskScheduler() {
	var wg sync.WaitGroup
	var taskList = []struct {
		Name string
		Fn   func()
	}{
		{"Task1", func() { time.Sleep(100 * time.Millisecond) }},
		{"Task2", func() { time.Sleep(200 * time.Millisecond) }},
		{"Task3", func() { time.Sleep(300 * time.Millisecond) }},
	}

	var taskStatList = make([]TaskStat, len(taskList))

	for i, task := range taskList {
		wg.Add(1)
		var startTime = time.Now()
		func(idx int, name string) {
			defer wg.Done()
			task.Fn() // 执行任务
			taskStatList[idx] = TaskStat{
				name, time.Since(startTime), true,
			}
		}(i, task.Name)
	}

	wg.Wait()
	fmt.Println("任务执行统计：")
	for _, taskStat := range taskStatList {
		fmt.Printf("%-8s 耗时: %6v\n", taskStat.Name, taskStat.Duration)

	}
}
