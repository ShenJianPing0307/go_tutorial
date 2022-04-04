package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 声明一个UserInfo的结构体
type UserInfo struct {
	UserName string
	Age      int
}

// 声明一个UserInfo的切片
type UserInfoSlice []UserInfo

// 实现interface Len Less Swap
func (uis UserInfoSlice) Len() int {
	// 返回切片的长度
	return len(uis)
}

func (uis UserInfoSlice) Less(i, j int) bool {
	// 按照什么样的顺序进行排序
	return uis[i].Age > uis[j].Age
}

func (uis UserInfoSlice) Swap(i, j int) {
	// 交换，更简单的交换方式 uis[i], uis[j] = uis[j], uis[i]
	temp := uis[i]
	uis[i] = uis[j]
	uis[j] = temp

}

func main() {
	// 随机生成不同年龄的UserInfo切片
	var uis UserInfoSlice
	for i := 0; i < 20; i++ {
		ui := UserInfo{
			UserName: fmt.Sprintf("NO~%d", rand.Intn(50)),
			Age:      rand.Intn(100),
		}

		// 将ui加入到uis中
		uis = append(uis, ui)
	}

	// 排序前顺序
	for _, v := range uis {
		fmt.Println(v)
	}

	// 调用sort.Sort
	sort.Sort(uis)

	// 排序后顺序
	fmt.Println()
	for _, v := range uis {
		fmt.Println(v)
	}

}
