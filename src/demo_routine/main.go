package main

import (
	"fmt"
)

func add(x, y int) {
	fmt.Println(x + y)
}

func count(ch chan int) {
	ch <- 1 // 向channel 写入数据，读取之前，操作是阻塞的

	fmt.Println("counting")

}
func main() {
	fmt.Println("hello")

	for index := 0; index < 10; index++ {
		go add(1, 2) // main函数不会等待goroutine，直接结束，所以不一定输出10次
	}

	chs := make([]chan int, 10) // 定义10个channel
	for index := 0; index < 10; index++ {
		chs[index] = make(chan int)
		go count(chs[index])
	}
	for _, ch := range chs {
		<-ch // 读出数据 ，操作也是阻塞的
	}

	// var name chan  elementType   ,var 变量名  chan  传输类型 。
	var chanint chan int      // 声明一个传输int的channel
	var _ map[string]chan int // map key:string  value: int channel

	var _ = make(chan int)
	chanint <- 1      // 将一个数据发送到chan ，向channel写入数据通常会导致阻塞，直到其他goroutine从这个channel中读取数据
	close(chanint)    // 关闭channel
	var a = <-chanint // 从channel 中读取数据
	fmt.Println(a)

	select {
	case <-chanint:
		fmt.Println(<-chanint)
	case chanint <- 2:
		fmt.Println("写入数据")
	default:
		fmt.Println("default")
	}

	var _ chan int   // 双向
	var _ chan<- int // 单向写
	var _ <-chan int // 单向读
	/* 所有的代码应该遵循，最小权限原则
	从而避免没有必要的使用泛滥问题，导致程序失控
	单向channel起到一种契约作用 */
	_, ok := <-chanint
	if ok == false {
		fmt.Println("chanint 已经关闭")
	}
	// fatal error: all goroutines are asleep - deadlock!  发生死锁
}
