package main

import "fmt"

func main() {
	dog := animal{
		name:  "dog",
		age:   1,
		color: "black",
	}
	dog.eat("apple")

	f := &animal{
		name:  "cat",
		age:   2,
		color: "yello",
	}
	f.eat("orange")

}

type animal struct {
	name  string
	age   int
	color string
}

func (a *animal) eat(food string) {
	fmt.Println(a.name, "eat ", food)
}

type food interface {
	eat(name string)
}
