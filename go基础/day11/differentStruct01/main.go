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

	// 第二个结构体变量，声明成指针变量
	var u2 *User = &u1
	fmt.Println(*u2) // {alily 20}

	u2.Name = "tomi"
	fmt.Println(u1, *u2) // {tomi 20} {tomi 20}

	// u1的地址、u2的地址、u2的值
	fmt.Printf("u1的地址：%p，u2的地址：%p，u2的值：%p", &u1, &u2, u2) // u1的地址：0xc000004078，u2的地址：0xc000006030，u2的值：0xc000004078
}
