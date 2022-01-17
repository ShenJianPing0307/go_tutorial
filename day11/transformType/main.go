package main

import "fmt"

type A struct {
	n1 string
}

type B struct {
	n1 string
}

func main() {
	// 类型转换
	var a A
	var b B
	a = A(b)

	fmt.Println(a, b)

}
