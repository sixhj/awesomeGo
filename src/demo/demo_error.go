package main

/* 2019-09-20 liu

go语言没有异常处理，一般通过返回值处理错误。
你也可以自己创建一个错误类型；唯一的要求就是它必须实现内置类型error接口：

*/

import (
	"errors"
	"fmt"
	"io"
)

type error interface {
	Error() string
}

func testError(count int) error {
	if count < 1 {
		return errors.New("Invalid count")
	}

	return nil
}

var EOF = errors.New("EOF") //该变量是可以访问的（因为第一个字母是大写字母

func main() {
	var input int
	_, err := fmt.Scan(&input)
	if err == io.EOF {
		fmt.Println("no more input!")
	}
}

//go语言有panic和recover函数。panic类似抛出异常，而recover类似捕获异常；但是很少使用它们
