package main

import "fmt"

// Golang中，一个自定义类型将接口中所有的方法都实现了，才能说这个自定义类型实现了该接口
type Animal interface {
	eat()
}

type Cat struct {
}

/*
这里如果写成这样，就是错误的，因为是*Cat指针类型实现了接口，而非Cat类型实现接口
func (c *Cat) eat() {

}

*/
func (c Cat) eat() {

}

func main() {

	c := Cat{}
	var a Animal = c
	fmt.Println(a)

}
