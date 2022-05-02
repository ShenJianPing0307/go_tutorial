package main

import "fmt"

func main() {

	// 写法一
	// for i := 0; i < 10; i++ {

	// 	fmt.Println("i=",i)
	// }

	// 写法二
	// i := 0
	// for i < 10 {
	// 	fmt.Println("i=",i)
	// 	i++
	// }

	// 写法三
	i := 0
	for {
		if i < 10 {
			fmt.Println("i=", i)
		} else {
			break
		}
		i++
	}

}
