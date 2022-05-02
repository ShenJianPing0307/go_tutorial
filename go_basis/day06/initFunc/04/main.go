package main

import "fmt"

func main() {
	// 方式一
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(5, 10)

	fmt.Println(res)

}
