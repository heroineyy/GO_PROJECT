package models

import (
	"errors"
	"ginweb/database"
)

type Users struct {
	Id       int    `gorm:"column:id" json:"id" form:"id"`                                          //ID
	Password string `gorm:"column:password" json:"password" form:"password" validate:"required"`    //密码
	UserName string `gorm:"column:user_name" json:"user_name" form:"user_name" validate:"required"` //用户名
}

func GetUser(username string, password string) (bool, error) {
	sql := " select  1 from  users where user_name=? and password=? "
	var isExists int
	database.MysqlDB.Raw(sql, username, password).Row().Scan(&isExists)
	if isExists > 0 {
		return true, errors.New("登录成功")
	} else {
		return false, errors.New("登录失败,用户名或者密码错误！")
	}
}

func RegisterUser(username string, password string) (bool, error) {
	var isExists int
	database.MysqlDB.Raw(" select 1 from  users where user_name=?", username).Row().Scan(&isExists)
	if isExists > 0 {
		return false, errors.New("创建用户失败,用户名已存在")
	}
	user := Users{UserName: username, Password: password}
	result := database.MysqlDB.Create(&user)
	if result.Error != nil {
		return false, errors.New("创建失败")
	}
	return true, errors.New("创建用户成功")
}
