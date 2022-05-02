package main

import "fmt"

// 空接口interface{}没有任何方法，所有类型都实现了空接口，即可以把任何一个类型的变量赋值给空接口
type A interface {
}

func main() {
	var x int
	var i A = x
	fmt.Println(i)
}
