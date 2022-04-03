package main

import (
	"newbubble/dao"
	"newbubble/models"
	"newbubble/routers"
)

func main() {
	//创建数据库
	//sql:CREATE database bubble;
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() //程序退出关闭数据库连接
	//模型绑定
	models.InitModel()
	//注册路由
	r := routers.SetupRouter()

	r.Run(":8081")
}
