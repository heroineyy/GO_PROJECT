package main

import "fmt"

//创建一个USB接口
type Usb interface {
	Start()
	Stop()
}

//手机和手机实现工作停止方法
type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

//相机和相机实现工作停止方法
type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

//电脑
type Computer struct {
}

//编写一个工作方法，接受usb接口类型的变量
//只要是是实现了Usb接口（所谓实现Usb接口，就是指是实现了Usb接口声明所有方法）
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	//先创建一个电脑，手机，相机
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.Working(phone)
	computer.Working(camera)

}