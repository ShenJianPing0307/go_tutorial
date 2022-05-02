package main

import "fmt"

func main() {

	// 数组初始化
	// 方式一
	var numArr1 [5]int = [5]int{1, 2, 3, 4, 5}
	// 方式二
	var numArr2 = [5]int{1, 2, 3, 4, 5}
	// 方式三
	var numArr3 = [...]int{1, 2, 3, 4, 5}
	// 方式四
	var numArr4 = [...]int{1: 2, 0: 1, 2: 3, 3: 4, 4: 5}

	// 方式五 类型推导
	numArr5 := [...]string{"zhangsan", "lisi", "wangwu"}
	fmt.Println("numArr1", numArr1)
	fmt.Println("numArr2", numArr2)
	fmt.Println("numArr3", numArr3)
	fmt.Println("numArr4", numArr4)
	fmt.Println("numArr5", numArr5)

}

/*
numArr1 [1 2 3 4 5]
numArr2 [1 2 3 4 5]
numArr3 [1 2 3 4 5]
numArr4 [1 2 3 4 5]
numArr5 [zhangsan lisi wangwu]
*/
