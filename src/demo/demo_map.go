package main

import "fmt"

/* 2019-09-20 liu
映射也叫map

k v 形式，一堆键值对，未排序的集合。

映射是动态增长的。然后，我们也可以在使用make时传递第二个参数设置映射的初始大小：

lookup := make(map[string]int, 100)

编译映射并不是有序的。每次遍历映射时，返回的键值对都是随机的顺序。

*/

type PersonInfo struct {
	Id      int
	Name    string
	Address string
	Age     int
}

func main() {

	var map1 map[int]string //  map[keyType]  valueType  声明map

	map1 = map[int]string{1: "map"} // 初始化

	map1[2] = "two"

	fmt.Println(map1)

	strings := make(map[int]string) // 使用make 初始化一个map

	strings[0] = "a"
	strings[1] = "ab"
	strings[13] = "abc" // 元素赋值
	for key, value := range strings {
		fmt.Println(key, value)
	}
	delete(strings, 0) // 删除元素
	//funcName1()

	// 查找一个map 中的值
	value, ok := strings[13]
	if ok {
		fmt.Println("map  有key = 13的值 = ", value)
	}
}

func funcName1() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]
	// 打印：0和false
	// 0代表一个整数型的默认值
	fmt.Println(power, exists)
	total := len(lookup)
	// 获取长度
	fmt.Println(total)
	lookup = map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}
}
