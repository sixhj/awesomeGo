package main

import "fmt"

func main() {
	f := squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) //4
	fmt.Println(f()) // 9

}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
