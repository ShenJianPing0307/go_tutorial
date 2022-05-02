package main

import "fmt"

func main() {
	// 求出三个数的最大值
	// 先求出两个数的最大值，然后再求出第三个

	a := 10
	b := 35
	c := 25

	var max int

	if a > b {
		max = a
	} else {
		max = b
	}

	fmt.Printf("两个数的最大值为%v", max) //35

	if c > max {
		max = c
	}

	fmt.Printf("三个数的最大值为%v", max) //35

}

// func main() {
// 	//交换a,b两个数的值，不允许使用中间变量
// 	a := 2
// 	b := 5

// 	a = a + b
// 	b = a - b // b = a + b - b
// 	a = a - b // a = a - (a-b)
// 	fmt.Printf("交换后a=%v,b=%v",a,b) //交换后a=5,b=2

// }

/*
func main() {
	// 交换a,b两个数的值
	a := 2
	b := 5
	fmt.Printf("交换前a=%v,b=%v",a,b) //交换前a=2,b=5
    //定义一个中间变量temp
	temp := a
	a = b
	b = temp
	fmt.Printf("交换后a=%v,b=%v",a,b) //交换后a=5,b=2

}
*/
