package main

import "fmt"

func main() {

	var s4 []string = []string{"zhangsan", "lisi", "wangwu"} 

	for i := 0; i < len(s4); i++ {

		fmt.Println(i, s4[i])

	}

}