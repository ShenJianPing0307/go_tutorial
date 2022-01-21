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
}

// Camera结构体变量接口实现方法
func (camera Camera) start() {
	fmt.Println("相机开始工作...")
}

func (camera Camera) stop() {
	fmt.Println("相机结束工作...")
}

// 声明一个Computer结构体
type Computer struct {
}

// usb会自动根据传递过来的来判断是phone还是camera然后调用对应的结构体变量中实现的方法
func (computer Computer) use(usb Usb) {
	usb.start()
	usb.stop()
}

func main() {
	// 创建结构体变量
	computer := Computer{}

	phone := Phone{}
	camera := Camera{}

	computer.use(phone)
	computer.use(camera)

}

/*
手机开始工作...
手机结束工作...
相机开始工作...
相机结束工作...
*/
