package main

import "fmt"

type Student struct {
	username string
	age      int
	int      // 匿名字段（基本数据类型）
}

func main() {
	// 匿名字段
	s := Student{}
	s.username = "peter"
	s.age = 20
	s.int = 85
	fmt.Println(s)
}
