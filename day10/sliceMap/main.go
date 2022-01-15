package main

import "fmt"

func main()  {
	
	// 定义切片
	class := make([]map[string]string, 2) //可以放入2个学生信息

	// 增加第一个学生
	if class[0] == nil {
		class[0] = make(map[string]string)
		class[0]["name"] = "张三"
		class[0]["gender"] = "男"
	}

	// 增加第二个学生
	if class[1] == nil {
		class[1] = make(map[string]string)
		class[1]["name"] = "李四"
		class[1]["gender"] = "男"
	}

	// 此时只能动态添加map了，否则会越界
	newStu1 := map[string]string{
		"name": "王五",
		"gender": "男",
	}

	class = append(class, newStu1)

	fmt.Println(class)
	
	/*
	[map[gender:男 name:张三] map[gender:男 name:李四] map[gender:男 name:王五]]
	*/

}