package main

import "fmt"

type integer int

// integer类型方法
func (i *integer) change() {
	*i = *i + 5
}

func main() {
	var i integer = 5
	i.change()
	fmt.Println(i) // 10
}
