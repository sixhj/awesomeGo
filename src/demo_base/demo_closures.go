package main

import "fmt"

// 返回的 内部的匿名函数
func intSeq() func() int { // 1
	i := 0
	return func() int { //2  5
		i++ //6
		return i
	}
}
func main() {
	fmt.Println("闭包")
	nextInt := intSeq() // 3

	fmt.Println(nextInt()) //4    执行的是内部的匿名函数
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
