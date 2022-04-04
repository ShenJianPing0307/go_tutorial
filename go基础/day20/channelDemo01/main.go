package main

import "fmt"

func main() {
	// 创建一个可以存放5个int类型的管道
	var intChan chan int = make(chan int, 5)

	// intChan是一个引用类型
	fmt.Printf("intChan的值=%v, intChan本身的地址=%p \n", intChan, &intChan) // intChan的值=0xc000020090, intChan本身的地址=0xc000006028

	// 向管道中写入数据,写入的数据不能超过容量
	intChan <- 20
	intChan <- 10

	// 查看管道的长度和容量（cap）
	fmt.Printf("len=%v, cap=%v \n", len(intChan), cap(intChan)) // len=2, cap=5

	// 读取管道中的值
	var num1 int
	num1 = <-intChan
	fmt.Println(num1) // 20

	//再次读取长度和容量
	fmt.Printf("len=%v, cap=%v \n", len(intChan), cap(intChan)) // len=1, cap=5
}
