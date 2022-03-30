package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func main() {
}

func doBadthing(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func timeout(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

//利用 time.After 启动了一个异步的定时器，返回一个 channel，当超过指定的时间后，该 channel 将会接受到信号。
//使用 select 阻塞等待 done 或 time.After 的信息，若超时，则返回错误，若没有超时，则返回 nil。

func test(t *testing.T, f func(chan bool)) { //车是是否关闭了管道
	t.Helper()
	for i := 0; i < 1000; i++ {
		timeout(f)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func TestBadTimeout(t *testing.T) { test(t, doBadthing) }
