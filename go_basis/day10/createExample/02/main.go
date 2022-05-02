package main

import "fmt"

func main() {

	// 创建map的方式二
	m2 := make(map[string]string)
	m2["usernamae"] = "张三"
	m2["password"] = "abc123456"
	fmt.Println(m2)

}
