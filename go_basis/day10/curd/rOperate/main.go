package main

import "fmt"

func main() {

	user := map[string]string{
		"username": "iveBoy",
		"gender":   "male",
	}

	val, isFlag := user["username"]

	fmt.Println(val, isFlag) // iveBoy true
}
