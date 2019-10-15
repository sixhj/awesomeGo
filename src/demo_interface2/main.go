package main

import "fmt"

type Eat interface {
	eat(food string)
}

type Drink interface {
	drink(water string)
}

type Live interface { // 接口组合
	Eat
	Drink
}

type Dog struct {
	name string
	age  int
}

type Cat struct {
	name string
}

func (d *Dog) eat(food string) {
	fmt.Println(d.name, " eat ", food)
}
func (d *Cat) eat(food string) {
	fmt.Println(d.name, " eat ", food)
}
func main() {
	// 接口
	var l Eat
	fmt.Printf("%T", l) //<nil>
	fmt.Println("")
	l = &Dog{
		name: "xiao gou ",
	}
	l.eat("bone")
	fmt.Printf("%T", l) //*main.Dogxiao
	fmt.Println("")
	l = &Cat{
		name: "xiao mao",
	}
	l.eat("fish")
	fmt.Printf("%T", l) //*main.Cat
	fmt.Println("")
	live := &Dog{
		name: "zhangsan",
	}
	live.eat("apple")

	fmt.Printf("%T", live) // *main.Dog

	var v1 interface{} = 1 // 任何对象实例都满足空接口，可以指向任何对象的any类型
	fmt.Println(v1)

}
