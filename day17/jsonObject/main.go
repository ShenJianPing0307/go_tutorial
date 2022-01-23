package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体
type User struct {
	Name string
	Age  int
}

func main() {

	// 创建结构体变量
	user := User{
		Name: "lily",
		Age:  20,
	}

	// 1、对结构体变量进行序列化
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("序列化错误 %v", err)
		return
	}
	// data是字节切片需要转成字符串
	str := string(data)
	fmt.Println(str)

	// 2、反序列化
	// 声明一个User结构体变量
	var user1 User
	err = json.Unmarshal([]byte(str), &user1)
	if err != nil {
		fmt.Printf("反序列化错误 %v", err)
	}
	fmt.Println(user1)
}
