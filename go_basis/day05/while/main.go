package main

import "fmt"

func main() {
	i := 0

	for {
		fmt.Println("i=", i)
		i++
		if i > 10 {
			break
		}

	}

}
