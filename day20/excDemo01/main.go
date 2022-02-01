package main

import "fmt"

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
	}
	defer close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, flag := <-intChan
		if !flag {
			break
		}
		fmt.Println(v)
	}
	// 完成读取协程，向exitChan发送信息
	exitChan <- true
	// 关闭管道
	defer close(exitChan)
}

func main() {
	var intChan chan int = make(chan int, 50)
	var exitChan chan bool = make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, flag := <-exitChan
		if flag {
			// 说明读取协程结束,可以结束主协程了
			break
		}
	}

}
