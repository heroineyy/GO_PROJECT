package main

import (
	"chatroom/client/process"
	"fmt"
)

var userId int
var userPwd string
var userName string

func main() {

	//接受用户的选择
	var key int
	//判断是否还继续显示菜单
	for true {
		fmt.Println("----------欢迎登录yy的多人聊天室----------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t  请选择（1-3）:")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的ID")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanln(&userPwd)
			//完成登录
			//1、创建一个UserProcess的实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户的ID")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanln(&userPwd)
			fmt.Println("请输入用户的名字")
			fmt.Scanln(&userName)
			//调用UserProcess,完成注册的请求
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
		default:
			fmt.Println("你输入有误，请重新输入")
		}
	}

	////更新用户的输入，显示新的提示信息
	//if key == 1 {
	//	//说明用户要登录
	//
	//	//先把登录的函数，写到另外一个文件，比如login.go
	//	//这里我们会需要重新调用
	//	///login(userId, userPwd)
	//
	//} else if key == 2 {
	//	fmt.Println("进行用户注册的逻辑。。。")
	//}

}
