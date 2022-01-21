package main

import (
	"fmt"
	"sort"
)

func main() {
	// 声明一个元素是int类型的切片
	var intSlice = []int{2, 1, 10, 5}

	// 对切片进行排序
	sort.Ints(intSlice)

	// 打印排序好的切片
	fmt.Println(intSlice)
}
