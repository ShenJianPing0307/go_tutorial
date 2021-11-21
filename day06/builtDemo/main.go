package main

import "fmt"

func main() {
	// len的使用

	var str string = "hello"
	fmt.Println("str长度", len(str)) // str长度 5

	// new的使用
	number := new(int) // 创建一个number变量，类型为*int，值为系统分配的地址
	*number = 10
	fmt.Printf("number的类型为%T，number的值为%v",number,number) // number的类型为*int，number的值为0xc000014098

}