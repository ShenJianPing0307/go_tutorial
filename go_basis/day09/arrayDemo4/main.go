package main

import "fmt"

func main() {

	var intArr [3]int = [3]int{4, 5, 6}

	for i, v := range intArr {
		fmt.Printf("i=%v, v=%v \n", i, v)
	}

}

/*
i=0, v=4
i=1, v=5
i=2, v=6
*/
