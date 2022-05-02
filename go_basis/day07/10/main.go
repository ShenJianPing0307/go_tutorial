package main

import (
	"fmt"
	"strings"
)

func main() {

	ifFlag := strings.EqualFold("ABC", "abc")
	fmt.Printf("ifFlag=%v", ifFlag) // ifFlag=true

}
