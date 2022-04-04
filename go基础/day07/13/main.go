package main

import (
	"fmt"
	"strings"
 )

func main() {

	index := strings.LastIndex("ABC", "C")
	fmt.Printf("str index=%v", index) // str index=2
	
}