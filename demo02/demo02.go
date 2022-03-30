package main

import (
	"fmt"
	"time"
)

//协程间的通信，管道。
func main() {
	//其负责协程之间的通信

	/**创建管道的三种方式
	var ch1 chan string
	ch1 = make(chan string)
	*/
	//ch1 := make(chan interface{})
	//
	///**通道操作符
	//<-
	//*/
	//
	////向管道中加入数据
	//ch1 <- 1 //用通道 ch 发送变量（双目运算符，中缀 = 发送）
	//
	////从通道流出（接收）
	//str := <-ch1
	//fmt.Println(str)

	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(time.Second)
}
func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}
func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
