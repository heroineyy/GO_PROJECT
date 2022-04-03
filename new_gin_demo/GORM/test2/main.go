package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//定义模型
type User struct {
	gorm.Model
	//ID   int64 //默认为组件
	Name string
	Age  int64
}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	////3.创建
	//u := User{Name: "XYY", Age: 99} //在代码层面创建一个User对象
	//fmt.Println(db.NewRecord(&u))   //判断主键是否为空
	//db.Debug().Create(&u)           //在数据库中创建了一条xyy的记录
	//fmt.Println(db.NewRecord(&u))

	//4.查询
	//var user User//声明模型结构类型变量user
	user := new(User) //new和make的区别
	db.First(&user)
	fmt.Println("user:%#v\n", user)

	var users []User
	db.Debug().Find(&users) //查询所有记录
	fmt.Println("user:%#v\n", user)

}
