package main

import (
	"fmt"
	"strings"
)

func main() {

	num := strings.Count("abcdef", "abc")
	fmt.Printf("num=%v", num) // num=1

}
