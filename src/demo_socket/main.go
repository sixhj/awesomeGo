package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("hello")
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	fmt.Println("tcp", conn, err)
	// conn.Read() 接受数据
	// conn.Write() 发送数据

}
