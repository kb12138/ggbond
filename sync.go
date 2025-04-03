package main

import (
	"fmt"
	"sync"
	"time"
)

// TaskWorker 是一个用于管理和执行任务的结构体
type TaskWorker[T any] struct {
	nWorkers    int
	processFunc func(T) // 自定义任务处理函数
}

// NewTaskWorker 创建一个新的 TaskWorker 实例，接受自定义的任务处理函数
func NewTaskWorker[T any](nWorkers int, processFunc func(T)) *TaskWorker[T] {
	return &TaskWorker[T]{
		nWorkers:    nWorkers,
		processFunc: processFunc,
	}
}

// Run 启动任务处理，接受一个任务列表
func (tw *TaskWorker[T]) Run(tasks []T) {
	tasksCh := make(chan T, len(tasks))
	var wg sync.WaitGroup

	// 启动指定数量的工作者 goroutines
	for i := 0; i < tw.nWorkers; i++ {
		wg.Add(1)
		go tw.worker(tasksCh, &wg)
	}

	// 将任务发送到任务通道
	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh) // 发送完毕后关闭通道

	// 等待所有工作者完成
	wg.Wait()
}

// worker 是一个内部方法，用于处理任务
func (tw *TaskWorker[T]) worker(tasksCh <-chan T, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasksCh {
		tw.processFunc(task) // 使用用户提供的处理函数
	}
}

// 自定义任务处理函数
func customProcess(task int) {
	fmt.Printf("自定义处理任务 %d\n", task)
	time.Sleep(time.Second) // 模拟任务处理时间
}
