package main

import "fmt"

//协程的同步：关闭通道-测试阻塞的通道
func main() {
	//通道可以被显式的关闭；尽管它们和文件不同：不必每次都关闭。
	//只有在当需要告诉接收者不会再提供新的值的时候，才需要关闭通道。只有发送者需要关闭通道，接收者永远不会需要。
	ch := make(chan interface{})
	defer close(ch)
	if v, ok := <-ch; ok {
		justiceClose(ch)
		fmt.Println(v)
	}
}
func justiceClose(ch chan interface{}) bool { //如何来检测可以收到没有被阻塞（或者通道没有被关闭）？
	v, ok := <-ch // ok is true if v received value
	fmt.Println(v)
	return ok
}
