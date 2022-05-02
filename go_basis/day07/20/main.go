package main

import (
	"fmt"
	"strings"
)

func main() {

	IsPrefix := strings.HasPrefix("hello world", "hello")
	fmt.Printf("IsPrefix=%v\n", IsPrefix) // IsPrefix=true

}
