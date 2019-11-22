package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("demo")

	var i int = 1
	fmt.Println(i)  //1
	fmt.Println(&i) // 0xc000090010

	pi := &i
	fmt.Println(pi)                 //0xc000090010  是变量的地址
	fmt.Println(&pi)                //0xc000096008
	fmt.Println(*pi)                // 1   = *(&i)  指针指向 地址变量的时候，取该地址变量的值
	fmt.Println(*(&i))              // 1
	fmt.Println(reflect.TypeOf(pi)) // *int

	//&  获取该变量的地址
	//* 获取该地址的值

	arr := []int{1, 2, 5, 2, 7, 1, 8, 4, 6, 8}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println(arr)

	sum := 1000000
	for i := 0; i < sum; i++ {
		wg.Add(1)
		go say(i)
		fmt.Println("启动线程个数", i)
	}
}

func say(s int) {
	fmt.Println(s, "线程启动")
	sum := 0
	for i := 0; i < s; i++ {
		sum += i
	}
	time.Sleep(time.Minute * 11)
	wg.Done()
}

func testScan() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Printf(input.Text())
	}

}
