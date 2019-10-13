package main

import "fmt"

func main() {

	point1 := Point{1, 2}
	point2 := Point{11, 22}

	fmt.Println(Add(point1, point2))
	fmt.Println(point1.Add(point2)) // 直接调用 该类型的方法

	fmt.Println(point2.sum()) // 编译器会对变量进行&point1 的隐式转换
	fmt.Println((&point2).sum())
}

type Point struct {
	X, Y int
}

// 普通的函数
func Add(p1, p2 Point) (int, int) {
	return p1.X + p2.X, p1.Y + p2.Y
}

// Point 类型的方法
func (p Point) Add(p2 Point) (int, int) {
	return p.X + p2.X, p.Y + p2.Y
}

func (p *Point) sum() int {
	return p.Y + p.X
}
