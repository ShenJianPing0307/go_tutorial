package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (user User) getUser() {

	fmt.Println(user.Name, user.Age)

}

func main() {

	var user User
	user.Name = "bily"
	user.getUser()

}
