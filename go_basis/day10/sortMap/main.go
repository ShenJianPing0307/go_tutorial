package main

import (
	"fmt"
	"sort"
)

func main() {

	m1 := make(map[int]int, 5)
	m1[1] = 5
	m1[3] = 10
	m1[0] = 60

	// 定义一个切片
	var keys []int

	for k, _ := range m1 {
		keys = append(keys, k)
	}

	// 对key进行排序
	sort.Ints(keys)

	fmt.Println(keys)

	// 根据key输出value
	for _, k := range keys {
		fmt.Println("value=", m1[k])
	}

}
