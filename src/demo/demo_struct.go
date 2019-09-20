package main

import "fmt"

type Apple struct {
	name string
	age  int
}

func main() {

	var a = Apple{name: "liu", age: 0}
	fmt.Println(a)

	b := Apple{name: "h", age: 2}
	fmt.Println(b)

	//testVar()

	addAge(a) // +1  失败  改变的只是一个拷贝的对象，不是原始的值
	fmt.Println("addAge\t\t\t", a)

	// 传入的是a结构体的地址
	// & 取地址  符
	aa := &a

	addAge2(aa) // + 2
	fmt.Println("addAge2\t\t\t", a)

	// 结构体上的函数，通过结构体的地址来调用
	aa.addAge3() // +3
	fmt.Println("addAge3\t\t\t", a)

	apple2 := getApple("j", 1)
	fmt.Println(apple2)

	/* 2019-09-19 liu
	创建结构体的两种方法
	*/
	a3 := new(Apple)
	a3.age = 33
	fmt.Println("a3.age = ", a3.age) //a3.age =  33
	fmt.Println(&a3)                 //0xc000082008
	fmt.Println(*a3)                 //{ 33}
	// 等效
	a4 := &Apple{}
	a4.age = 44
	fmt.Println("a4.age = ", a4.age) //a4.age =  44
	fmt.Println(a4)                  //&{ 44}
	fmt.Println(*a4)                 //{ 44}  //获取指针指向的值
	fmt.Println(&a4)                 //0xc00000e020

}

func testVar() {
	var i = 1
	// 类型推断
	p := 2
	// 声明 并 赋值
	var s int
	// 声明 默认值是0
	fmt.Println(i, p, s)
}

func addAge(a Apple) Apple {
	a.age += 1
	return a
}

// * 代表执行apple 的指针
func addAge2(s *Apple) {
	s.age += 2
}

// 将结构体 和 方法关联
func (s *Apple) addAge3() {
	s.age += 3
}

func getApple(n string, a int) Apple {
	return Apple{
		name: n,
		age:  a,
	}
}

/////////////

type Driver struct {
	name string
}
type Car struct {
	d       *Driver //  Driver 类型的指针
	*Driver         // 当一个字段是指针类型没有给定 字段名的时候，我们可以间接使用这个字段和方法，go在编译的时候会给该字段一个名字，
}

func getCar() {
	car := Car{
		d:      &Driver{name: "d"},
		Driver: &Driver{name: "d2"},
	}
	fmt.Println(car)

}
