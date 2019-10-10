package main // 是一个特殊的package ， main 函数是一个特殊的函数，是一个程序的入口

import "fmt" // 输入输出、格式化等

/* 2019-09-19 liu
多行注释
一行代表一个语句结尾，不需要分隔符
多行语句写在一行，可以用分号隔开

命名规则：
a-z A-Z 0-9  下划线_

类型：
布尔、整型、字符串、派生类型


派生类型：指针类型（Pointer）、数组、结构化（struct）、Channel、函数、切片、接口、Map

支持 goto

*/

var g = 1 //  全局变量
func main() {
	fmt.Println("Hello  GO !")
	var age int
	age = 1
	fmt.Print(age)

	name := "liu" // 简写 声明变量 ， 根据值自动判断类型， 只能在函数中使用
	fmt.Print(name)

	//var b bool = true // 布尔类型

	//	 int  float   未运算，采用补码
	var c, d, e int
	fmt.Print(" int e = ", e)

	c = 2
	d = 3
	fmt.Print(c + d)

	//_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值

	//const HEIGHT int   = 2    常量
	// iota，特殊常量，可以认为是一个可以被编译器修改的常量。

	/* 定义结构体 */
	type Circle struct {
		radius float64
	}

	var arr [11]int
	arr[1] = 0
	var arr2 = []int{1, 2, 3}
	arr2[1] = 1

	var a = 20  /* 声明实际变量 */
	var ip *int /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Println("ip 变量储存的指针地址: %x\n", ip)
	/* 2019-09-19 liu
	   当一个指针被定义后没有分配到任何变量时，它的值为 nil。

	   nil 指针也称为空指针。

	   nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

	   一个指针变量通常缩写为 ptr。
	*/

	//testStruct()

	//testMap()

	//testFor()

	//testIf()

}

func testMap() {
	var countryCapitalMap map[string]string // [key] value
	/*创建集合 */
	countryCapitalMap = make(map[string]string)
	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"
}

func testStruct() {
	type Apple struct {
		name    string
		size    int
		color   string
		address string
	}
}

func testIf() {
	if 1 > 2 {
		fmt.Println("???")
	} else if 2 > 1 {

	} else {
		fmt.Println(1)
	}
}

func testSwitch() {
	switch i {
	case 1:
		fmt.Println(1)
		//fallthrough   会执行下一个case
	case 2:
		fmt.Println(2)

	default:
		fmt.Println(3)
	}

}

func testFor() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// func 方法名（参数。。。 类型）返回值类型 {
func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// a, b := swap("Google", "Runoob")

func swap(x, y string) (string, string) {
	return y, x
}

/* 定义接口 */
type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}
