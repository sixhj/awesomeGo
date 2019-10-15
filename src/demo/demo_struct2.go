package main

import "fmt"

func main() {

	var num int = 9
	fmt.Printf("num address=%v\n", &num)

	var ptr *int
	ptr = &num // prt
	*ptr = 10  //这里修改时，会到num的值变化
	fmt.Println("num =", num)
	// var a_b int = 20
	// fmt.Println(a_b)

	var int int = 30 // int 是内建类型 ，可以当作变量名，可以重新定义
	fmt.Println("int 是内建类类型，可以作为变量名，可以重新被定义", int)

	apple2 := &apple{ // & 可以调用apple 的所有方法（包括*apple）
		// apple 只能调用 apple的方法
		name: "my apple",
	}
	apple2.eat()

	apple3 := &apple{ // 必须加
		name: "one apple",
	}
	apple3.eat()
	eat(apple3)
}

type apple struct {
	name string
}

func eat(a *apple) { // 文件内的私有方法， 方法的参数必须是一个地址
	fmt.Println("eat apple ", a.name)
}

func (a *apple) eat() {
	fmt.Println("eat apple2", a.name)
}

/*
T的方法集仅拥有 T Receiver （方法中的接受者）方法。
*T 方法集则包含全部方法 (T + *T)。
*/
