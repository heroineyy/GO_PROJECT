package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

var (
	MysqlDB  *gorm.DB
	username = "root"
	password = "root"
	dbName   = "user"
	hostname = "127.0.0.1"
	port     = 3306
)

func init() {
	var err error
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, hostname, port, dbName)
	MysqlDB, err = gorm.Open("mysql", connArgs)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("连接接成功")
	MysqlDB.SingularTable(true) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	MysqlDB.LogMode(true)       //打印sql语句
	//开启连接池
	MysqlDB.DB().SetMaxIdleConns(100)   //最大空闲连接
	MysqlDB.DB().SetMaxOpenConns(100)   //最大连接数
	MysqlDB.DB().SetConnMaxLifetime(30) //最大生存时间(s)
}
