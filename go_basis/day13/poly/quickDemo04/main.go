package main

import "fmt"

// 一个自定义类型可以实现多个接口
type A interface {
	a1()
}

type B interface {
	b1()
}

type integer int

func (i integer) a1() {

}

func (i integer) b1() {

}

func main() {
	var i integer = 10
	var a A = i
	var b B = i
	fmt.Println(a, b)
}
