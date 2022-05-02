package main

import "fmt"

func main() {

	// 创建map的方式一 先进行make
	var m1 map[string]string
	m1 = make(map[string]string, 10)
	m1["usernamae"] = "张三"
	m1["password"] = "abc123456"
	fmt.Println(m1)

}
