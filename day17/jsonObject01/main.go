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

	// 定义一个切片
	var userSlice []User

	// 将结构体变量加入到切片中
	userSlice = append(userSlice, user)
	// 1、对切片进行序列化
	data, err := json.Marshal(userSlice)
	if err != nil {
		fmt.Printf("序列化错误 %v", err)
		return
	}
	// data是字节切片需要转成字符串
	str := string(data)
	fmt.Println(str)

	// 2、反序列化
	// 声明一个待反序列化的切片
	var userSlice1 []User
	err = json.Unmarshal([]byte(str), &userSlice1)
	if err != nil {
		fmt.Printf("反序列化错误 %v", err)
	}
	fmt.Println(userSlice1)
}
