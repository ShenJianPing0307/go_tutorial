package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	// 创建结构体变量
	var user User
	fmt.Println(user) // { 0} 分别是空字符串和0
}
