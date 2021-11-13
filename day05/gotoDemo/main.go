package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i > 5 {
			goto label1
		}
		fmt.Println("i=", i)
	}
	label1:
	fmt.Println("goto1")
	fmt.Println("goto2")

}
/*
输出：
i= 0
i= 1
i= 2
i= 3
i= 4
i= 5
goto1
goto2

*/