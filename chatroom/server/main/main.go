package main

import (
	"chatroom/server/model"
	"fmt"
	"net"
	"time"
)

//处理和客户端的通讯
func process(conn net.Conn) {
	//读客户端发送的信息
	defer conn.Close()
	//这里调用总控，创建一个
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端通信协程错误=err", err)
		return
	}
}

//完成对UserDao的初始化实例
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	//当服务器启动时，初始化redis连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
}

func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听。。。")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//一旦连接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
