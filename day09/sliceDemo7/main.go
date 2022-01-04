package main

import "fmt"

func main() {

	str7 := "abcdef" // 在内存中以字节数组的形式存在，[6]byte

	// 使用切片获取到abc
	s7 := str7[0:3]

	fmt.Println(s7) // abc

}