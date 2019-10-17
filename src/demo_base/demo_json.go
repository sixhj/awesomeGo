package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Movie struct {
		Title string `json:"title"`
	}

	var ms = []Movie{
		{Title: "go"}}
	bytes2, e := json.Marshal(ms) // 数组-》json
	fmt.Println(string(bytes2), e)

	// 字符串解析为结构体
	s := `{"id": 1, "name": "wxnacy"}`

	type User struct {
		Id   int
		Name string
	}
	var user User
	// 将字符串反解析为结构体
	json.Unmarshal([]byte(s), &user)
	fmt.Println(user) // {1 wxnacy}

	var d map[string]interface{}
	// 将字符串反解析为字典
	json.Unmarshal([]byte(s), &d) // json -》 map
	fmt.Println(d)                // map[id:1 name:wxnacy]

}
