package main

import (
	"fmt"
	"time"
)

func main() {
	//什么是协程：协程是轻量的，比线程更轻。它们痕迹非常不明显（使用少量的内存和资源）：使用 4K 的栈内存就可以在堆中创建它们。
	//因为创建非常廉价，必要的时候可以轻松创建并运行大量的协程（在同一个地址空间中 100,000 个连续的协程）。
	//并且它们对栈进行了分割，从而动态的增加（或缩减）内存的使用；栈的管理是自动的，但不是由垃圾回收器管理的，而是在协程退出后自动释放。
	//var numCores = flag.Int("n", 2, "number of CPU cores to use")
	//flag.Parse()
	//runtime.GOMAXPROCS(*numCores) //必须设置 GOMAXPROCS 为一个大于默认值 1 的数值来允许运行时支持使用多于 1 个的操作系统线程

	fmt.Println("In main()")
	go longWait()  //开启一个longWait协程
	go shortWait() //开启一个shortWait协程
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}
func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}

//main()，longWait() 和 shortWait() 三个函数作为独立的处理单元按顺序启动，然后开始并行运行。
//每一个函数都在运行的开始和结束阶段输出了消息。
//为了模拟他们运算的时间消耗，我们使用了 time 包中的 Sleep 函数。
//Sleep() 可以按照指定的时间来暂停函数或协程的执行，这里使用了纳秒（ns，符号 1e9 表示 1 乘 10 的 9 次方，e=指数）。
