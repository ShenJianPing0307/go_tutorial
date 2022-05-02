package main

import "fmt"

func modifyValue(m map[int]int) {
	m[1] = 10
}

func main() {

	m2 := map[int]int{
		1: 6,
		2: 9,
	}

	// 修改map值
	modifyValue(m2)

	fmt.Println(m2)
	/* map[1:10 2:9] */
}
