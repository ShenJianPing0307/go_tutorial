package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	// 创建一个结构体变量
	user := User{
		"kity",
		23,
	}
	fmt.Println(user)

	// 将结构体变量序列化
	jsonStr, err := json.Marshal(user)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(string(jsonStr)) // {"name":"kity","age":23}
	}

}
