package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 基本使用
	var a = false
	fmt.Println("a=", a) //a= false

	// 占用字节
	fmt.Println("a占用的字节大小为", unsafe.Sizeof(a)) //a占用的字节大小为 1

}
