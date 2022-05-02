package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var intArr [5]int

	len := len(intArr)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		intArr[i] = rand.Intn(100) // 生成 大于等于0，小于100的5个数
	}

	fmt.Println("交换前：", intArr)

	// 进行交换
	temp := 0
	for i := 0; i < len/2; i++ {
		temp = intArr[len-1-i]
		intArr[len-1-i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println("交换后：", intArr)

}

/*
交换前： [42 43 95 0 67]
交换后： [67 0 95 43 42]
*/
