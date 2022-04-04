package main

import "fmt"

type animal struct {
	name  string
	hobby string
}

func (a *animal) showHobby() {
	fmt.Println("hobby:", a.hobby)
}

type cat struct {
	animal
	name string
}

func (c *cat) showHobby() {
	fmt.Println("hobby:", c.hobby)
}

type tigger struct {
	animal
	cat
}

func main() {
	/*
		1、结构体可以使用匿名结构体中的所有字段和方法，包括首字母小写的结构体、属性、方法等私有变量
		2、匿名结构体字段访问可简化
		3、当结构体于匿名结构体有相同字段和方法时，采用就近原则，如若希望访问匿名结构体字段需要指明匿名结构体
	*/
	// 声明一个cat结构体变量
	var c cat
	c.name = "猫"      // 就近原则，使用的name时cat结构体中的name，而非animal中的name
	c.hobby = "吃鱼..." // 匿名结构体可简写，查找顺序：本结构体-->各个匿名结构体-->如若没有报错
	c.showHobby()

	/*
		4、如若结构体有两个或者多个匿名结构体时，并且匿名结构体都拥有相同字段和方法，但是结构体本身无同名字段和方法时，需指明匿名结构体，否则编译报错
	*/
	// 声明一个tigger结构体
	t := tigger{}
	t.animal.name = "老虎..."    // 需要指明具体的匿名结构体
	t.animal.hobby = "吃小动物..." // 需要指明具体的匿名结构体
	t.animal.showHobby()

}
