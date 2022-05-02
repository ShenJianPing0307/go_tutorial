package main

import "fmt"

func main() {

	// 声明一个map

	user := map[string]string{

		"username": "iveBoy",
		"gender":   "male",
	}

	for k, v := range user {
		fmt.Printf("k=%v v=%v\n", k, v)
	}
	/*
		k=username v=iveBoy
		k=gender v=male
	*/

}
