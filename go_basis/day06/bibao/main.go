package main

import "fmt"

// 累加器
func AddUpper() func(int) int {

	var n int = 5
	return func(x int) int {
		n = n + x
		return n
	}

}

func main() {

	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))

}

/*
输出：
6
8
*/
