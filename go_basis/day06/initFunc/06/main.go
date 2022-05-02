package main

import "fmt"

var (
	Func = func(n1 int, n2 int) int {
		return n1 + n2
	}
)

func main() {
	// 全局匿名函数调用
	res := Func(10, 5)
	fmt.Println(res)

}
