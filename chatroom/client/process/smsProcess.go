package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//创建一个SmesMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err.Error())
		return
	}
	mes.Data = string(data)

	//对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err.Error())
		return
	}
	//5、将mes发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	//6、发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err.Error())
		return
	}
	return

}
