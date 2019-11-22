package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	maps = make(map[int]uint64, 10)
)

var wg sync.WaitGroup

var lock sync.Mutex

func getFactorial(num int) {
	defer wg.Done()
	res := 1

	for i := 1; i < num; i++ {
		res *= num
	}

	lock.Lock()
	maps[num] = uint64(res)
	lock.Unlock()
}
func main() {
	fmt.Println("cpu number is ：", runtime.NumCPU())
	wg.Add(22)
	for i := 0; i < 22; i++ {
		go getFactorial(i)
	}
	wg.Wait()
	for key, value := range maps {
		fmt.Println("key : ", key, "\t value : ", value)
	}
}
