package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func send(numChan chan<- int) {
	// chan<- 只写
	defer wg.Done()
	for i := 0; i < 20; i++ {
		numChan <- i
	}
	close(numChan)

}

func recv(numChan <-chan int) {
	// <-chan 只读
	defer wg.Done()
	for {
		v, flag := <-numChan
		if !flag {
			break
		}
		fmt.Println(v)
	}
}

func main() {

	var numChan chan int = make(chan int, 20)
	wg.Add(2)

	go send(numChan)
	go recv(numChan)

	wg.Wait()

}
