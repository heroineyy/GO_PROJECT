package main

import (
	"chatroom/common/message"
	"chatroom/server/process"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

//先创建一个Processor的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个serverProcessMes函数
//功能：根据客户端发送信息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录登录
		//创建一个Processor的结构体实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegester(mes)
	case message.SmsMesType:
		//创建一个SmsProcess实例完成转发群聊消息
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (this *Processor) process2() (err error) {
	//循环的客户端发送信息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回message,err
		//创建一个Transfer实例完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err != io.EOF {
				fmt.Println("客户端退出，服务器端也退出。。")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
