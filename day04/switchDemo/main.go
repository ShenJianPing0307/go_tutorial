package main

import "fmt"

func main() {
	
	switch score := 90; {
		case score >= 90:
			fmt.Println("优秀...")
			fallthrough

		case score >= 80 && score < 90:
			fmt.Println("良好...")
			
		default:
			fmt.Println("一般...")
	}


}