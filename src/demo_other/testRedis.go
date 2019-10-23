package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
)

func main() {

}

func testRedis() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	fmt.Println("hello redis")
	_, err = c.Do("SET", "name", "six")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	username, err := redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get username: %v \n", username)
	}
	defer c.Close()
}

func funcName1() {
	var num int = 1
	fmt.Printf("num type = %T", num)
	fmt.Println("转义字符")
	/*
	 \t

	*/
	arr := [2]int{1, 2}
	fmt.Println(reflect.TypeOf(arr))
	sarr := []int{1, 2, 3}
	fmt.Println(reflect.TypeOf(sarr))
	ints := append(sarr, 7)
	//追加
	fmt.Println(ints)
	i := copy(ints, sarr)
	//3  返回实际上复制的元素个数
	fmt.Println(i, len(ints))
	sarr[1] = 11
	ints[2] = 22
	for _, value := range ints {
		fmt.Print(value, ",")
	}
	fmt.Println()
	for key, value := range sarr {
		fmt.Println(key, "=", value)
	}
	i2 := make(map[string]int)
	i2["a"] = 123
}
