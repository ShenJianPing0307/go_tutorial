package main

import "fmt"

func main() {

	var s4 []string = []string{"zhangsan", "lisi", "wangwu"} 

	for i, v := range s4 {
		fmt.Println(i, v)
	}

}