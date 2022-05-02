package main

import "fmt"

func main() {

	var str string = "hello 中国"
	str1 := []rune(str)

	//for-range方式

	// for index, value := range str {
	// 	fmt.Printf("index=%d value=%c \n", index, value)
	// }

	// 传统方式
	for i := 0; i < len(str1); i++ {

		fmt.Printf("%c \n", str[i])

	}

}
