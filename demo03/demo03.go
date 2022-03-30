package main

import (
	"fmt"
	"time"
)

//通道阻塞

func main() {
	//对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的：
	//如果ch中的数据无人接收，就无法再给通道传入其他数据：新的输入无法在通道非空的情况下传入。
	//所以发送操作会等待 ch 再次变为可用状态：就是通道值被接收时（可以传入变量）。

	ch1 := make(chan int)
	//go pump(ch1)       // pump hangs
	//fmt.Println(<-ch1) // prints only 0

	go pump(ch1)
	go suck(ch1)
	time.Sleep(time.Second)

}
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
