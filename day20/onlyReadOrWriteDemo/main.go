package main

import "fmt"

func main() {

	var numChan chan<- int // 声明一个只写管道

	numChan <- 1

	var numChan1 <-chan int // 声明一个只读管道
	var num1 = <-numChan1
	fmt.Println(num1)

}
