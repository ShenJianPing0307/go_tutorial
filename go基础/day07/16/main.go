package main

import (
	"fmt"
	"strings"
 )

func main() {

	str := strings.TrimSpace(" Hello World ")
	fmt.Printf("str=%v", str) // str=Hello World
	
}
/*
str1=hello world
str2=HELLO WORLD
*/