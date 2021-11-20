package main

import "fmt"

func main() {

	a := func(n1 int, n2 int) int {
		return n1 + n2
	}

	res := a(10, 5)
	fmt.Println(res)
	
}