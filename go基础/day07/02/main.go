package main

import "fmt"

func main() {

	var str string = "hello北京" 
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("i=%c\n",str2[i])
	}
	
}
/*
i=h
i=e
i=l
i=l
i=o
i=北
i=京

*/