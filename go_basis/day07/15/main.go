package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Hello World"
	str1 := strings.ToLower(str)
	str2 := strings.ToUpper(str)
	fmt.Printf("str1=%v\n", str1)
	fmt.Printf("str2=%v", str2)

}

/*
str1=hello world
str2=HELLO WORLD
*/
