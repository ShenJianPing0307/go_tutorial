package main

import (
	"fmt"
	"strings"
 )

func main() {

	str := strings.TrimLeft("? hello world?", "?")
	fmt.Printf("str=%q\n", str) // str=" hello world?"
	
}