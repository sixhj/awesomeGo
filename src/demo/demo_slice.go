package main

import "fmt"

func main() {
	//声明切片可以指定 【】 大小
	scores := []int{1, 4, 293, 4, 9}

	// 使用make 时需要指定  make（类型，长度，容量）
	/* 2019-09-20 liu
	创建一个切片不仅仅是分配一段内存（new只能分配一段内存），
	需要为底层数组分配内存，并且也要初始化这个切片
	长度表示切片的长度，容量表示底层数组的大小
	*/
	scores2 := make([]int, 0, 10)
	fmt.Println(scores, scores2)

	//testNewSlice(scores)

	//testAppend(scores)
}

func testNewSlice(scores []int) {
	//	 初始化切片的 方法
	names := []string{"leto", "jessica", "paul"}
	//得提前知道你想往数组存放的值
	checks := make([]bool, 10)
	//往切片的特定位置写入一个值时
	var names []string
	//返回一个空切片，一般和append一起使用，此时切片的元素数量是未知的
	scores := make([]int, 0, 20)
	//指定切片的初始容量。当我们大概知道需要多少元素时很有用。即使你知道元素的个数，append也能被使用
}

func testAppend(scores []int) {
	//	append会重新创建一个更大的数组，并将所有的值复制过去
	//	go扩展数组使用的是2倍算法
	scores := make([]int, 5)
	scores = append(scores, 9332)
	fmt.Println(scores)
	//[0, 0, 0, 0, 0, 9332]
}
