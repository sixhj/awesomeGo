package main

import "fmt"

func main() {
	var scores [10]int
	scores[0] = 339
	fmt.Println(scores[0])
	fmt.Println(scores)
	fmt.Println("hello")

	scores2 := [4]int{9001, 9333, 212, 33}
	for index, value := range scores2 {
		fmt.Println(index, value)
	}

	/* 2019-09-19 liu
	切片是一个轻量级的结构体封装，这个结构体被封装后，代表一个数组的一部分

	*/
	scores3 := []int{1, 4, 293, 4, 9}
	fmt.Println(scores3)

	scores4 := make([]int, 10) //创建一个切片不仅仅是分配一段内存（new只能分配一段内存）
	fmt.Println(scores4)

	//初始化了一个长度和容量都是10的切片。长度表示切片的长度，容量表示底层数组的大小。
	//scores := make([]int, 0, 10) //长度为0但是容量为10的切片

}
