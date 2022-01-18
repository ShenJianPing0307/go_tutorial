package main

import "fmt"

type CalSum struct {
	n1 int
	n2 int
}

// 结构体调用方法，该方法可以有入参参数、返回值
func (cal *CalSum) getSum() int {
	res := cal.n1 + cal.n2
	return res
}

func main() {

	// 生成一个结构体变量
	cal := CalSum{
		1,
		2,
	}
	/*
		调用方法, 因为方法的形参是指针类型，所以应该是结构体变量地址调用
	*/

	res := (&cal).getSum() // .的运算优先级高，所以(&cal)需要带括号
	fmt.Println(res)       // 3

	/*
		但是编译器在底层进行了优化，所以(&cal).getSum() 可以写成cal.getSum()
	*/

	res1 := cal.getSum()
	fmt.Println(res1)
	fmt.Printf("%p", &cal)

	/*
		此时main中的结构体变量与方法中传入的结构体变量相同，内存中存在形式也与值传递不同
	*/

}
