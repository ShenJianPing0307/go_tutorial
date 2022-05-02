package main

import "fmt"

func main() {

	var s2 []int = make([]int, 3, 10)
	s2[0] = 1
	s2[1] = 2

	fmt.Println(s2) // [1 2 0]

}
