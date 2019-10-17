package main

import (
	"fmt"
	"reflect"
)

/* 2019-10-12 liu

 */

var c = 1 // 全局变量

var (
	d = 1
	s = 2
)

func main() {
	var x int
	x = 1

	fmt.Println(x)

	var s string

	s = "hello "
	fmt.Println(s + s)

	//fmt.Println(x + s)   类型不一致不能 相加

	var a = "11"                           // 类型推导
	fmt.Println(reflect.TypeOf(a), "类型推导") // string

	aa := 22
	fmt.Println(aa, "声明并赋值，同时类型推导")

	var age int
	age = 1
	fmt.Println(age)

	var rr rune // uncode  int32
	fmt.Println(rr)

}
