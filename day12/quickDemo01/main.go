package main

import "fmt"

type CalSum struct {
	n1 int
	n2 int
}

// 结构体调用方法，该方法可以有入参参数、返回值
func (cal CalSum) getSum() int {
	res := cal.n1 + cal.n2
	return res
}

func main() {

	// 生成一个结构体变量
	cal := CalSum{
		1,
		2,
	}
	// 调用方法
	res := cal.getSum()
	fmt.Println(res)
}
