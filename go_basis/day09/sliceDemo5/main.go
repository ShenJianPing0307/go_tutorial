package main

import "fmt"

func main() {

	// 声明一个切片
	var s5 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(s5) // [1 2 3 4 5]

	// 通过append给切片s5追加元素
	s5 = append(s5, 6, 7, 8, 9)
	fmt.Println(s5) // [1 2 3 4 5 6 7 8 9]

	// 通过append追加切片给s5
	s5 = append(s5, s5...)
	fmt.Println(s5) // [1 2 3 4 5 6 7 8 9 1 2 3 4 5 6 7 8 9]

}
