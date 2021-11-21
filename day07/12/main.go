package main

import (
	"fmt"
	"strings"
 )

func main() {

	str := strings.Replace("ABA", "Z", "Z", -1)
	fmt.Printf("str=%v", str) // str=ABA
	
}