package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printGoroutine() {
	// 协程执行函数
	defer wg.Done() // 每执行完1次该协程，Add中的参数减1
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Hello goroutine %v \n", i)
	}

}

func main() {
	// 主协程中运行main函数
	// 开启一个协程
	wg.Add(1) // Add中的参数是开启的协程数量
	go printGoroutine()

	// 继续执行主协程代码
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Hello main %v \n", i)
	}
	wg.Wait() // 阻塞，直到Add中的参数变为0结束该主协程

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
Hello goroutine 4
*/
