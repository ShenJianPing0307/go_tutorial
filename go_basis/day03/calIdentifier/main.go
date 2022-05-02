package main

import "fmt"

func main() {
	// + - * / % ++ --案例

	var a int = 12
	var b int = 7

	fmt.Println(a + b) // 19
	fmt.Println(a - b) // 5
	fmt.Println(a * b) // 84
	fmt.Println(a / b) // 1 如果需要为浮点数，重新定义a为浮点数12.0 var a float32 = 12.0
	fmt.Println(a % b) // 5
	a++
	fmt.Println(a) // 13
	a--
	fmt.Println(a) //12

}
