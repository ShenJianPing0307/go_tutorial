package main

import "fmt"

// 接口本身不能创建实例，但是可以指向一个实现了该接口的自定义数据类型的实例
type A interface {
	m1()
}

type S1 struct {
}

func (s1 S1) m1() {
	// 实现A接口中的m1方法
}

func main() {
	s1 := S1{}
	var a1 A = s1
	fmt.Println(a1)
}
