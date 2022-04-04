package main

import "fmt"

func main() {

	var A bool = true
	var B bool = false

	fmt.Println(A&&B) //false
	fmt.Println(A||B) // true
	fmt.Println(!(A&&B)) //true
}