package main

import "fmt"

func main() {

	var intChan chan int = make(chan int, 1000)

	for i := 0; i < 800; i++ {

		intChan <- i // 管道中放入800个数

	}
	// 放入数字后，关闭管道，循环结束不会报deadlock问题
	close(intChan)

	// 遍历管道使用for-range, 不能使用普通循环
	for v := range intChan {
		fmt.Println("v=", v)
	}

}
