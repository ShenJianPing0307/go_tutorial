package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 方式一
	var p1 Person
	fmt.Println(p1)
}
