package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var numChan chan int = make(chan int, 1000)

func write(numChan chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		numChan <- i
	}
	//close(numChan) //下面通过select解决了如果不关闭管道出现阻塞而产生deadlock问题

}

func read(numChan chan int) {
	defer wg.Done()

	for {
		// 有时并不清楚何时关闭管道，通过select解决未关闭管道出现deadlock问题
		select {
		case v := <-numChan:
			fmt.Println(v)
			time.Sleep(time.Second)

		default:
			fmt.Println("读取结束")
			return

		}

	}

}

func main() {
	wg.Add(2)
	// 启动2个协程
	go write(numChan)
	go read(numChan)

	wg.Wait()

}
