package main

import "fmt"

func main() {
	// 此时变量i的作用域就是在代码块中
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 如果需要使用这个i变量可以定义成局部变量
	var i int
	for i = 0; i < 5; i++ {
		fmt.Println(i)
	}



}