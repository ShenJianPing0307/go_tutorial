package main

import "fmt"

func main() {
	// 打印小于等于5的数

	for i := 0; i <= 10; i++ {
		if i > 5 {
			break
		} else {
			fmt.Println("i=", i)
		}
	}

}
