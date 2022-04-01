# Channel
Channel learing note

# demo

主要是简单的介绍和sync使用。

# demo01

介绍协程使用 4K 的栈内存就可以在堆中创建它们。

因为创建非常廉价，必要的时候可以轻松创建并运行大量的协程（在同一个地址空间中 100,000 个连续的协程）。

并且它们对栈进行了分割，从而动态的增加（或缩减）内存的使用；栈的管理是自动的，但不是由垃圾回收器管理的，而是在协程退出后自动释放。

运行结果：

```In main()

About to sleep in main()

Beginning shortWait()

Beginning longWait()

End of shortWait()

End of longWait()

At the end of main()
```


# 协程之间的通信demo02

创建管道的三种方式：
```
var ch1 chan string
ch1 = make(chan string)
ch1 := make(chan interface{})
```

输出结果：
```
Washington Tripoli London Beijing Tokyo
```

