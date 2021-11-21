package main

import (
	"fmt"
	"strconv"
 )

func main() {

	number, error := strconv.Atoi("123")
	if error != nil {
		fmt.Println("error=", error)
	} else {
		fmt.Println("number=", number) // number= 123
	}
	

}