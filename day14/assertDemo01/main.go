package main

import "fmt"

func main() {

	var i interface{}

	var a float64 = 6.32
	i = a //空接口可以接收任意数据类型
	// 使用类型断言
	b := i.(float64)
	fmt.Printf("b的类型为%T 值为%v", b, b) // b的类型为float64 值为6.32

}
