package main

import (
	"fmt"
	"go_tutorial/day06/initFunc/03/utils" //引入utils包
)

var age = test()

func test() int {
	fmt.Println("test...")
	return 12
}

func init() {
	fmt.Println("init...")

}

func main() {
	fmt.Println("main...")
	fmt.Println("init...", utils.Age)

}

/* 输出
utils包init...
test...
init...
main...
init... 20
*/
