package main

import "fmt"

func main() {

	var intArr [3]int = [3]int{4, 5, 6}

	fmt.Println("第三个元素", intArr[2]) // 第三个元素 3

	for i := 0; i < len(intArr); i++ {
		fmt.Printf("第%d个元素=%d \n", i, intArr[i])
	}

}

/*
第0个元素=4
第1个元素=5
第2个元素=6
*/
