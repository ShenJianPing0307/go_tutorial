package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i > 5 && i < 9 {
			continue
		}
		fmt.Println("i=", i)
	}
}
/*
输出：
i= 0
i= 1
i= 2
i= 3
i= 4
i= 5
i= 9
*/