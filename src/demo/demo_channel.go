package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* 2019-09-20 liu
同步

使用channel

通道通过解决数据共享问题，让并发编程变得更加清晰。
通道是一个通信管道，它用于go协程之间传递数据。
换句话说，go协程可以通过通道，传递数据给另外一个go协程。
其结果就是，在任何时候，仅有一个go协程可以访问数据。
通道与所有其他的东西一样，也有类型。这个类型，就是将要在通道中传递的数据的类型。


从一个通道接收或者发送数据时会阻塞。
当我们从一个通道接收数据时，直到数据可用, go协程才会继续执行。
类似的，往一个通道发送数据时，在数据被接收之前, go协程也不会继续执行。

go语言确保，往一个通道发送数据时，仅有一个单独的接收器可以接收。


通道是唯一共享的状态，通过通道，可以安全的，并发发送和接收数据。
通道提供了我们需要的所有同步代码，并且也确保，在任意的特定时刻，

只有一个go协程，可以访问数据的特定部分。



*/

func main() {
	c := make(chan int)

	//CHANNEL <- DATA  //往通道发送数据
	//VAR := <-CHANNEL  //从通道接收数据

	//使用select的最主要目的是，通过它管理多个通道。
	// 给定多个通道，select将阻塞直到有一个通道可用。
	// 如果没有可用的通道，当提供了default语句时，执行该分支。
	// 当多个通道都可用时，选择其中的一个通道是随机的。
	for {
		select {
		case c <- rand.Int():
			//可选的代码
		default:
			//这里可以留下空行以丢弃数据
			fmt.Println("dropped")
		}
		time.Sleep(time.Millisecond * 50)
	}
}