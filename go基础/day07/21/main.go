package main

import (
	"fmt"
	"strings"
 )

func main() {

	IsSuffix := strings.HasSuffix("hello world", "world")
	fmt.Printf("IsSuffix=%v\n", IsSuffix) // IsSuffix=true
	
}