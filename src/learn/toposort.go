package main

import "fmt"

func main() {
	// 反映了所有课程和先决课程的关系

	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calulus":    {"linear algebra", "test"},

		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
	}

	fmt.Print(prereqs)

	fmt.Println()

	type duck struct {
		age  int
		Bird // 内嵌 构成复杂的类型
	}

	i := new(duck)
	i.eat = false // 直接访问内部的字段，不需要声明
	i.name = "six "
	i.fly() // 访问内嵌类型的方法

}

//func (p Point) Add(p2 Point) (int, int) {
type Bird struct {
	eat  bool
	name string
}

func (b *Bird) fly() {
	fmt.Println(b.name, "flying...")
}
