package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// 定义一个map类型
	var userMap map[string]interface{}

	//map使用前需要进行make
	userMap = make(map[string]interface{}, 5)

	userMap["UserName"] = "Alice"
	userMap["Age"] = 18

	// 1、对切片进行序列化
	data, err := json.Marshal(userMap)
	if err != nil {
		fmt.Printf("序列化错误 %v", err)
		return
	}
	// data是字节切片需要转成字符串
	str := string(data)
	fmt.Println(str)

	// 2、反序列化
	// 声明一个待反序列化的map
	var userMap1 map[string]interface{}
	err = json.Unmarshal([]byte(str), &userMap1)
	if err != nil {
		fmt.Printf("反序列化错误 %v", err)
	}
	fmt.Println(userMap1)
}
