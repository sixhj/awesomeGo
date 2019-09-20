package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(" os demo")
	name, _ := os.Hostname()
	fmt.Println(name)
	s, _ := os.UserHomeDir()
	fmt.Println(s)

	env := os.Environ()
	for k, v := range env {
		fmt.Println(k, v)
	}

	dir, _ := os.Getwd()
	fmt.Println(dir)
}
