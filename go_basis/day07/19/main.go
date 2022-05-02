package main

import (
	"fmt"
	"strings"
)

func main() {

	str := strings.TrimRight("? hello world?", "?")
	fmt.Printf("str=%q\n", str) // str="? hello world"

}
