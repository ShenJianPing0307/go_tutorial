package main

import "fmt"

type A struct {
	x1 string
	x2 int
}

type B struct {
	x1 float64
}

type C struct {
	A
	B
}

func main() {
	c := C{}
	// 必须指明匿名结构体的
	c.A.x1 = "abc"

	fmt.Println(c)
}
