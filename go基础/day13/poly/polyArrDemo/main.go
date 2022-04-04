package main

import "fmt"

// 定义一个接口
type Usb interface {
	// 声明两个未实现的方法
	start()
	stop()
}

// 声明一个Phone的结构体
type Phone struct {
	Name string
}

// Phone结构体变量接口实现方法
func (phone Phone) start() {
	fmt.Println("手机开始工作...")
}

func (phone Phone) stop() {
	fmt.Println("手机结束工作...")
}

// 声明一个Camera的结构体
type Camera struct {
	Name string
}

// Camera结构体变量接口实现方法
func (camera Camera) start() {
	fmt.Println("相机开始工作...")
}

func (camera Camera) stop() {
	fmt.Println("相机结束工作...")
}

func main() {
	// 创建Usb结构体数组变量
	var usbArr [3]Usb
	usbArr[0] = Phone{"iphone11"}
	usbArr[1] = Phone{"xiaomi"}
	usbArr[2] = Camera{"Canon"}
	fmt.Println(usbArr)

}

/*
[{iphone11}
{xiaomi}
{Canon}]
*/
