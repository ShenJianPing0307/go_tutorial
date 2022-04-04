package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	// 创建第一个结构体变量
	var u1 User = User{
		"alily",
		20,
	}
	fmt.Println(u1) // {alily 20}

	// 第二个结构体变量,结构体默认是值拷贝，所以u1与u2是两个不同的数据空间
	u2 := u1
	fmt.Println(u2) // {alily 20}

	u2.Name = "tomi"
	fmt.Println(u1, u2) // {alily 20} {tomi 20}
}
