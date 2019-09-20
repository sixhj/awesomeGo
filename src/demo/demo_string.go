package main

import "fmt"

func main() {
	stra := "the spice must flow"
	byts := []byte(stra)
	strb := string(byts)
	fmt.Println(byts)
	fmt.Println(strb)

	print(stra)

	for key, value := range byts {
		fmt.Println(key, value)
	}
	fmt.Println(len(byts))

	//当字符串有由unicode字符码runes组成时。如果你计算字符串的长度时，可能得到的结果和你期待的不同。

	print("asdf你好")
}

func print(stra string) {

	for key, value := range stra {
		fmt.Println(key, value)
	}
	fmt.Println(len(stra))
}

// 在你使用[]byte(X)或者string(X)时务必注意，你创建的是数据的拷贝。这是由于字符串的不可变性。

//如果你通过range遍历一个字符串，你将得到runes，而不是bytes。当然，当你将一个字符串转换成一个[]byte时，你将得到正确的数据。
