package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 方式三 var p1 *Person = new(Person)
	var p1 *Person = new(Person)
	/*
		此时p1是一个指针类型，如果赋值需要：
		(*p1).Name = "igon"
		(*p1).Age = 25
		但是，go设计者为了使用方便，已经在底层做了处理，所可以这样写：
		p1.Name = "igon"
		p1.Age = 25
	*/
	p1.Name = "igon"
	p1.Age = 25

	fmt.Print(*p1) // 指针取值 {igon 25}

}
