package main

import "fmt"

func main() {

	// 定义一个数组
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	// 定义一个切片，引用上面的数组
	var s1 = a[0:3]
	fmt.Println(s1) // [1 2 3]
	fmt.Println("切片的个数", len(s1)) // 片的个数 3
	fmt.Println("切片的动态容量", cap(s1)) // 切片的动态容量 5
	
}