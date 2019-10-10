package main

import "fmt"

func main() {
	//f1()
	s := [2]string{"a", "b"}
	for key, value := range s {
		fmt.Println(key, value)
	}

	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for _, v := range s {
		fmt.Println(v)
	}

	s[1] = "c"
	fmt.Println(s[1])
}

func f1() {
	var scores [10]int // 声明 长度为10的int型数组
	scores[0] = 339    //第0个元素赋值
	fmt.Println(scores[0])
	fmt.Println(scores)
	fmt.Println("hello")
	scores2 := [4]int{9001, 9333, 212, 33} // 声明长度4 的int 数组 同时赋值
	for index, value := range scores2 {    // 遍历数组
		fmt.Println(index, value)
	}
	/* 2019-09-19 liu
	切片是一个轻量级的结构体封装，这个结构体被封装后，代表一个数组的一部分
	*/
	scores3 := []int{1, 4, 293, 4, 9} // 不定长度的数组
	fmt.Println(scores3)
	scores4 := make([]int, 10)
	//创建一个切片不仅仅是分配一段内存（new只能分配一段内存）
	fmt.Println(scores4)
	//初始化了一个长度和容量都是10的切片。长度表示切片的长度，容量表示底层数组的大小。
	//scores := make([]int, 0, 10) //长度为0但是容量为10的切片
}

var a [1]int

//var a1 [1]*Apple

var a2 = [3]byte{1, 2}

var a3 = [1]string{"a"}

var a4 = [...]int{123, 1, 2, 3}
