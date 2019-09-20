package main

/* 2019-09-20 liu
接口是一种类型，它只定义了声明，没有具体实现
*/

type Logger interface {
	Log(message string)
}

//如果在编程时使用接口，而不是它们具体的实现，我们可以很容易的改变和测试我们代码，但是对我们的代码没有任何影响。

type Server struct {
	logger Logger
}

//接口常用于避免循环导入。由于接口没有具体的实现，所以他们的依赖性有限。

//最重要是的记住go语言中的包名和你的目录结构有密切关系（不仅仅在一个项目中，在整个工作空间都如此）。

//go拥有一个没有任何方法的空接口：interface{}。因为每种类型都实现了空接口的0个方法，并且接口都是隐式实现，所以每种类型都实现了空接口的条约

func addInterface(a interface{}, b interface{}) interface{} {

	return nil
}

//将一个空接口变量转换成一个指定的类型，你可以使用.(TYPE):
//return a.(int) + b.(int)
