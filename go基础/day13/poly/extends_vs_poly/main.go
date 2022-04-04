package main

import "fmt"

// 声明一个Monkey的结构体
type Monkey struct {
	Name string
}

// Monkey结构体天生拥有climbing方法
func (m *Monkey) Climbing() {
	fmt.Printf("%s爬树...", m.Name)
}

// 声明一个LittleMonkey结构体
type LittleMonkey struct {
	Monkey
}

// 声明一个接口
type OtherSkills interface {
	Swimming()
}

// LittleMonkey通过后天不断努力实现了OtherSkills
func (lk *LittleMonkey) Swimming() {
	fmt.Printf("%s学会了游泳...", lk.Name)
}

func main() {
	// 创建一个LittleMonkey的结构体变量
	lk := LittleMonkey{}
	lk.Name = "小猴子"

	// 调用天生的Climbing方法
	lk.Climbing()

	// 通过后天的努力实现了OtherSkills，学会了Swimming
	lk.Swimming()

}
