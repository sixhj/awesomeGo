package main

/* 2019-09-20 liu

垃圾回收

无论如何你的defer都会在方法返回时得到执行，虽然这有点极端。但是这可以帮你在初始化的附近释放资源，并且可以实现多个返回点。

你可以使用defer完成任何目的，例如当一个函数退出时打印日志记录。

*/

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("a_file_to_read")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 读这个文件
}
