package main

import (
	"fmt"
	"strings"
)

func main() {

	str := strings.Trim("? hello world? ", "?")
	fmt.Printf("str=%q\n", str) // str=" hello world? "

}
