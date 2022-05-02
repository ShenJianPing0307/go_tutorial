package main

import "fmt"

func main() {

	var totalLevel int = 20

	// i表示层数
	for i := 1; i <= totalLevel; i++ {
		// 打印前先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}
		// j 表示每层打印多少*
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()

	}

}
