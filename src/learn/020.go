package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num int = 1
	fmt.Printf("num type = %T", num)

	fmt.Println("转义字符")
	/*
	 \t

	*/

	arr := [2]int{1, 2}
	fmt.Println(reflect.TypeOf(arr))

	sarr := []int{1, 2, 3}
	fmt.Println(reflect.TypeOf(sarr))

	ints := append(sarr, 7) //追加
	fmt.Println(ints)

	i := copy(ints, sarr) //3  返回实际上复制的元素个数
	fmt.Println(i, len(ints))

	sarr[1] = 11
	ints[2] = 22
	for _, value := range ints {
		fmt.Print(value, ",")
	}

	fmt.Println()
	for key, value := range sarr {
		fmt.Println(key, "=", value)
	}

	i2 := make(map[string]int)
	i2["a"] = 123

}
