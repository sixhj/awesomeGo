package main

import "fmt"

// var name type  =  expression
// name := expression     name 的类型由 expression 决定
var age int = 1
var stra = "hello 你好"

var aa, b = 1, 2

var ( // var () 定义一组变量
	i int
	s string
	//f float32
)

func main() {

	ints := [...]int{1, 2, 3} // 1、整型数组
	ints = [3]int{1, 2, 3}    // 2、整型数组

	for key, value := range ints {
		fmt.Println("key = ", key, "value = ", value)
	}

	fmt.Println("--------1008")

	l := len(stra)
	fmt.Println("len(stra)", l)
	for i := 0; i < l; i++ {
		fmt.Println(i, stra[i])
	}

	fmt.Println("----1008----")

	i = 11
	s = "six"
	//f = 0.123

	const PI float32 = 3.14 // 常量

	const k = iota // 可以被编译器修改的常量 ， 每次增加1 ， 下一次 出现时 重置

	// _ 匿名变量

	//  :=  声明并初始化 ， 只能在函数内使用

	// 枚举 const () 定义一组常量
	const (
		Su = iota
		Mo
		Tu
		We
		Th
		Fr
		Sa
	)
	fmt.Println(Su)
	fmt.Println(Mo)
	fmt.Println(Sa)

	var ii int32 = 123
	i = int(ii) // int  是平台相关的 ， 需要做强转

	if Sa > Mo {
		fmt.Println("今天休息")
	}
	fmt.Println(byte(Sa))
	fmt.Println(Sa << 1)

	//nine()
	//testString()
}

func nine() {
	for i := 0; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", i*j, "\t")
		}
		fmt.Println()
	}
}

func testString() {
	stra := "the spice must flow"
	byts := []byte(stra) //字符串数组
	strb := string(byts)

	fmt.Println(byts)
	fmt.Println(strb)
	print(stra)
	for key, value := range byts {
		fmt.Println(key, value)
	}
	fmt.Println(len(byts))
	//当字符串有由unicode字符码runes组成时。如果你计算字符串的长度时，可能得到的结果和你期待的不同。
	print("asdf你好")
}

func printString(stra string) {

	for key, value := range stra {
		fmt.Println(key, value)
	}
	fmt.Println(len(stra))
}

// 在你使用[]byte(X)或者string(X)时务必注意，你创建的是数据的拷贝。这是由于字符串的不可变性。

//如果你通过range遍历一个字符串，你将得到runes，而不是bytes。当然，当你将一个字符串转换成一个[]byte时，你将得到正确的数据。
