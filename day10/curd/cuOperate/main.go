package main

import "fmt"

func main()  {
	
	user := map[string]string{
		"username": "iveBoy",
	}

	// 增加，如果key没有，就是增加
	user["password"] = "abc123456"

	// 修改，如果key存在，就是修改
	user["username"] = "lily"

	fmt.Println(user)

}