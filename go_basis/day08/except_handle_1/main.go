package main

import (
	"fmt"
)

func test() {

	defer func() {
		err := recover() // recover() 内置函数，可以捕获到异常
		if err != nil {
			fmt.Println("err=", err)
		}
	}()

	n1 := 5
	n2 := 0
	res := n1 / n2
	fmt.Println("res=", res)

}

func main() {
	test()
}

/*
err= runtime error: integer divide by zero
*/
