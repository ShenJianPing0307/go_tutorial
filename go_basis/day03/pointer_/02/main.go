package main

import "fmt"

func main() {

	// 错误的使用方法，指针存储的是内存地址，修改为var ptr *int = &a正确
	var int a = 10
	var ptr *int = a

	// 错误的使用方法，指针类型与值类型不匹配，修改为var ptr1 *int = &b正确
	var int b = 20
	var ptr1 *float32 = &b
}
