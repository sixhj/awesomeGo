package main

import "fmt"

/* 2019-09-20 liu
协程

go协程类似一个线程，但是go协程是由go自己调度，而不是系统。在协程中的代码可以和其他代码并发执行。

简单的将go关键字附在我们想要执行的函数前面即可

go协程很容易创建且开销较小。
最终多个go协程将会在同一个底层的系统线程上运行。
这也常称之为M:N线程模型，

因为我们有M个应用线程（go协程）运行在N个系统线程上。
结果就是，一个go协程的开销和系统线程比起来相对很低（一般都是几KB）。

在现代的硬件上，有可能拥有成千上万个go协程。

*/

import (
	"time"
)

func main() {
	fmt.Println("start")
	go process()
	time.Sleep(time.Millisecond * 10) // this is bad, don't do this!
	fmt.Println("done")

	// 可以使用匿名函数
	go func() {
		fmt.Println("processing")
	}()
}

func process() {
	fmt.Println("processing")
}
