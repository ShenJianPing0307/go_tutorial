package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体
type User struct {
	Name string `json:"name"` // tag别名 name
	Age  int    `json:"age"`  // tag别名 age
}

func main() {

	// 创建结构体变量
	user := User{
		Name: "lily",
		Age:  20,
	}

	// 对结构体变量进行序列化，可以看到输出的字段名是tag别名
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("序列化错误 %v", err)
		return
	}
	// data是字节切片需要转成字符串
	str := string(data)
	fmt.Println(str) // {"name":"lily","age":20}

}
