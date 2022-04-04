package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float64
}

func (student *Student) Examing() {
	fmt.Print("姓名:", student.Name, "年龄:", student.Age, "考试成绩:", student.Score)
}

type MiddleStudent struct {
	Student // 学生匿名结构体
}

type UgStudent struct {
	Student // 学生匿名结构体
}

// 大学生特有读书方法
func (ugs *UgStudent) ReadBook() {
	fmt.Print(ugs.Name, "正在读书...")
}

func main() {
	// MiddleStudent继承Student，所以可以通过mt.Student.Name来进行属性的调用，
	// 同时也可以mt.Name，如果在本结构体中没找到该属性，会去匿名结构体中查找
	mt := MiddleStudent{}
	mt.Student.Name = "中学生" //mt.Name = "中学生"
	mt.Student.Age = 15     //mt.Age = 15
	mt.Student.Score = 85.5 //mt.Score = 85.5
	mt.Examing()

	// 调用大学生特有的方法
	ugs := UgStudent{}
	ugs.Name = "大学生"
	ugs.Age = 25
	ugs.Score = 95.5
	ugs.ReadBook()

}
