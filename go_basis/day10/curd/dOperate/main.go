package main

import "fmt"

func main() {
	user := map[string]string{
		"username": "iveBoy",
		"gender":   "male",
	}

	// 删除键值
	delete(user, "gender")
	fmt.Println(user) // map[username:iveBoy]

	//  删除不存在的键, 删除不操作，也不会报错
	delete(user, "hobby")
	fmt.Println(user)
}
