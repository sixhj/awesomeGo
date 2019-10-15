package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")
	fmt.Println(os.Hostname())
	fmt.Println(os.UserHomeDir())
	fmt.Println("内存页大小", os.Getpagesize())
	env := os.Environ()
	for k, v := range env {
		fmt.Println(k, v)
	}
	os.FileMode
}
