package main

import (
	"fmt"
	"time"
)

func printGoroutine() {
	// 协程执行函数
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Hello goroutine %v \n", i)
	}

}

func main() {
	// 主协程中运行main函数
	// 开启一个协程
	go printGoroutine()

	// 继续执行主协程代码
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Hello main %v \n", i)
	}

}

/*
输出：
Hello main 0
Hello goroutine 0
Hello goroutine 1
Hello main 1
Hello main 2
Hello goroutine 2
Hello goroutine 3
Hello main 3
Hello main 4
*/
