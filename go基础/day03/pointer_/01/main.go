package main

import "fmt"

func main() {

	// 定义声明一个基本数据类型，其本身的内存地址 &i
	var a int = 10

	// 定义声明一个指针变量
	// 1、ptr是一个指针变量
	// 2、ptr的数据类型是*int
	// 3、prt本身的内存地址是 &ptr
	var ptr *int = &a

	fmt.Printf("ptr=%v\n", ptr) // ptr=0xc000014098
	fmt.Printf("ptr的地址是%v\n", &ptr) // ptr的地址是0xc000006028
	fmt.Printf("ptr指向的值是%v\n", *ptr) // ptr指向的值是10
}