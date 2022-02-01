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
	for i := 0; i < 500; i++ {
		numChan <- i
	}
	close(numChan) //写入完毕后一定要关闭管道，否则读取到最后会deadlock

}

func read(numChan chan int) {
	defer wg.Done()

	for {
		v, flag := <-numChan
		if !flag {
			break
		}
		fmt.Println(v)
		time.Sleep(time.Second)
	}

}

func main() {
	wg.Add(2)
	// 启动2个协程
	go write(numChan)
	go read(numChan)

	wg.Wait()

}
