package main

import (
	"fmt"
	"go_tutorial/day12/factoryMode01/model"
)

func main() {

	user := model.NewUser("lily", 20)
	fmt.Println(model.GetName(user))
	// fmt.Println(*user)               // {lily 20}
	// fmt.Println(user.Name, user.Age) // lily 20

}
