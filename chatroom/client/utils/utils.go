package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//这里将这些方法关联到结构体上
type Transfer struct {
	//分析他应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte //此刻传输使用缓冲
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据。。。")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("coon.Read err=", err)
		return
	}
	//fmt.Println("读到的buf=",buf[:4])
	//根据buf[:4]转成一个unit32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	//根据pkgLen 读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail err=", err)
		return
	}

	//把pkgLen 反序列化成 -> message.Mssage
	json.Unmarshal(this.Buf[:pkgLen], &mes) //一定要重视&
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}
