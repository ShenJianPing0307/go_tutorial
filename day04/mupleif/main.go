package main

import "fmt"

func main() {
    a := 10

	if a < 5 {
		fmt.Println("小于5")
	} else if a >= 5 && a < 10{
		fmt.Println("大于等于5并且小于10")
	} else {
		fmt.Println("大于10")

	}

}


