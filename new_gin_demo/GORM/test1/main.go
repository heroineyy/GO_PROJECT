package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     uint `gorm:"primary_key"`
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 自动迁移
	db.AutoMigrate(&UserInfo{}) //把结构体和表的字段进行绑定

	u1 := UserInfo{'5', "七米", "男", "篮球"}
	u2 := UserInfo{'7', "沙河娜扎", "女", "足球"}
	// 创建记录
	db.Create(u1)
	fmt.Println(db.NewRecord(&u1))
	db.Create(&u2)
	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	//db.Delete(&u)
}
