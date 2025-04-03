package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	////TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	//// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	//s := "gopher"
	//fmt.Println("Hello and welcome, %s!", s)
	//
	//for i := 0; i <= 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//for i := 1; i <= 5; i++ {
	//	//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
	//	// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
	//	fmt.Println("i =", 100/i)
	//}

	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 任务列表
	//nWorkers := 3  // 工作者数量
	var nWorkers = 3 // 工作者数量

	// 创建 TaskWorker 实例，传入自定义处理函数
	worker := NewTaskWorker(nWorkers, customProcess)

	// 运行任务，并提供任务列表
	worker.Run(tasks)

}
