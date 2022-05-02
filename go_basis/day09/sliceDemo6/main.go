package main

import "fmt"

func main() {

	// 声明一个切片
	var s6_1 []int = []int{1, 2, 3, 4, 5}

	// 初始化一个切片,默认整型的值都是0
	var s6_2 = make([]int, 10)

	// 将s6_1中的值拷贝到s6_2中
	copy(s6_2, s6_1)

	fmt.Println(s6_1) // [1 2 3 4 5]
	fmt.Println(s6_2) // [1 2 3 4 5 0 0 0 0 0]

}
