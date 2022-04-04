package main

import "fmt"

func main() {

	var i interface{}

	var a float64 = 6.32
	i = a //空接口可以接收任意数据类型
	// 使用类型断言，检测机制，避免panic错误
	b, flag := i.(float64)
	if flag {
		fmt.Printf("b的类型为%T 值为%v", b, b) // b的类型为float64 值为6.32
	} else {
		fmt.Println("类型转换失败")
	}

}
