package main

import "fmt"

// init函数，通常可以在init函数中完成初始化工作
func init() {
	fmt.Println("init...")
}

func main() {
	fmt.Println("main...")
}
/*
init...
main...
*/