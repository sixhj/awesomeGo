package main

import "fmt"

/* 2019-09-20 liu
映射也叫map

k v 形式

映射是动态增长的。然后，我们也可以在使用make时传递第二个参数设置映射的初始大小：

lookup := make(map[string]int, 100)

编译映射并不是有序的。每次遍历映射时，返回的键值对都是随机的顺序。

*/
func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]
	// 打印：0和false
	// 0代表一个整数型的默认值
	fmt.Println(power, exists)

	total := len(lookup) // 获取长度
	fmt.Println(total)

	lookup = map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}
}
