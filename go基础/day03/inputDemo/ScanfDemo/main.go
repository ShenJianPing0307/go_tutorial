package main

import "fmt"

func main() {
	var username string
	var password string

	fmt.Println("请输入用户名、密码，使用空格隔开")
	// 程序停止在此处，等待用户输入，然后回车
	fmt.Scanf("%s %s", &username, &password)
	fmt.Printf("用户名：%v，密码：%v",username, password)

}