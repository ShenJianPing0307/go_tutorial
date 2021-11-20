package main

import "fmt"

func test(number *int) {
	*number = 20 //修改number的值

}

func main() {
	number := 12
	test(&number) // 调用变量的值的函数
	fmt.Println("main...", number) // main... 20

}