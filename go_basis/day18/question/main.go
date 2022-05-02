package main

import "fmt"

// 被测试的函数
func calAdd(n1 int, n2 int) int {

	res := n1 + n2
	return res

}

func main() {
	// 传统测试方式，在main函数中调用，看输出结果
	res := calAdd(1, 2)
	if res != 3 {
		fmt.Printf("calAdd函数错误，期望值是%v，返回值是%v", 3, res)
	} else {
		fmt.Printf("calAdd函数正确，期望值是%v，返回值是%v", 3, res)

	}
}
