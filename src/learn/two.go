package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

}

func funcName() {
	// 复数
	x := complex(1, 2)
	fmt.Println(x)
	//(1+2i)
	y := complex(2, 3)
	fmt.Println(y)
	//(2+3i)
	fmt.Println(x * y)
	//(-4+7i)
	fmt.Println(real(x * y))
	//-4
	fmt.Println(imag(x * y))
	// 7
}

func osTest() {
	//funcName2()
	type Celsius float64
	type Fahrenheit float64
	const (
		AbsoluteZeroC Celsius = -273.15
		FreezingC     Celsius = 0
		BoilingC      Celsius = 100
	)
	dir, err := os.Getwd()
	if err == nil {
		fmt.Println(dir) ///Users/liuhuajian/go/src/awesomeProject
	}
	split := strings.Split(string(dir), "/")
	for key, value := range split {
		fmt.Println(key, value)
	}
	/* 2019-10-10 liu
	0
	1 Users
	2 liuhuajian
	3 go
	4 src
	5 awesomeProject
	*/
}

func funcName2() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", "", "separator")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	p := new(int)
	//  *int 型的p ，指向未命名的int变量
	*p = 2
	// 未命名的int 变量赋值为2
}
