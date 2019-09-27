package main

import "fmt"

/* 2019-09-19 liu

func [* Type] funcName(var1 ,var2 type) (type1,type2)

func (s *Apple) addAge3() { // 方法和结构体 结合

if语句中定义并初始化的值在if语句之外是不可用的，但是可以在else if和else语句中使用

支持多个返回值

导入包含该函数的包 就可以调用该函数了

小写字母的函数，在本包中可以用

大写字母的函数，才能被其他包调用

传入任意类型的参数 使用{}interface  例如： func Println(a ...interface{})

*/

/* 2019-09-21 liu

匿名函数
不带函数名的函数声明

函数可以想普通变量一样被传递或使用

go的匿名函数是一个闭包

- 概念
可以包含自由变量的代码块，这些变量不在这个代码块内或者任何全局上下文中定义，
而是在代码块环境中定义，
要执行的为自由变量提供绑定的计算环境（作用域）

闭包的价值在于可以作为函数对象或者匿名函数

这些函数可以存储到变量中作为参数传递给其他函数，最重要的是能够被函数动态创建和返回

闭包也可以应用外部变量，只要闭包一直存在，那么该外部变量就一直存在


*/

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sumProduct(2, 3))
	fmt.Println(sums(1, 2, 3))

}

func add(a, b int) int {
	return a + b
}

// 多个参数
func sumProduct(a, b int) (int, int) {
	return a + b, a * b
}

// 可变参数
func sums(numbers ...int) int {
	fmt.Println("numbers len : ", len(numbers))
	total := 0
	for pos, num := range numbers {
		fmt.Print(pos) // 索引
		total += num
	}

	fmt.Println()

	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i])
	}

	fmt.Println()
	for _, value := range numbers {
		fmt.Println("value = ", value)
	}
	return total

}
