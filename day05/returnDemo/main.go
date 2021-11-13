package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i > 5 {
			return
		}
		fmt.Println("i=", i)
	}
	fmt.Println("retunDemo...")

}

/*
输出：
i= 0
i= 1
i= 2
i= 3
i= 4
i= 5

*/