package main

import "fmt"

/* 2019-09-19 liu

func [* Type] funcName(var1 ,var2 type) (type1,type2)

func (s *Apple) addAge3() { // 方法和结构体 结合

*/

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sumProduct(2, 3))
	fmt.Println(sums(1, 2, 3))

}

func add(a, b int) int {
	return a + b
}

// 多个参数
func sumProduct(a, b int) (int, int) {
	return a + b, a * b
}

// 可变参数
func sums(numbers ...int) int {
	fmt.Println("numbers len : ", len(numbers))
	total := 0
	for pos, num := range numbers {
		fmt.Println(pos) // 索引
		total += num
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return total
}
