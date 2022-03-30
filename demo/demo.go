package main

import "sync"

type Info struct {
	mu sync.Mutex
	// ... other fields, e.g.: Str string
}

func main() {
	//什么是协程：协程是轻量的，比线程更轻。它们痕迹非常不明显（使用少量的内存和资源）：使用 4K 的栈内存就可以在堆中创建它们。
	//因为创建非常廉价，必要的时候可以轻松创建并运行大量的协程（在同一个地址空间中 100,000 个连续的协程）。
	//并且它们对栈进行了分割，从而动态的增加（或缩减）内存的使用；栈的管理是自动的，但不是由垃圾回收器管理的，而是在协程退出后自动释放。

}

func Updata(demo *Info) {
	//上锁
	demo.mu.Lock()
	//使用完释放锁
	demo.mu.Unlock()
}
