package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//服务器启动胡初始化一个userDao实例
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//1、根据用户ID返回一个User实例
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定ID去redis 查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err != nil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

//完成登录的校验Login
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从连接池取一个连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//这时证明这个用户获取到
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

//注册
func (this *UserDao) Register(user *message.User) (err error) {
	//先从连接池取一个连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err != nil {
		err = ERROR_USER_EXISTS
		return
	}
	//这时id在redis还没有，则完成注册
	data, err := json.Marshal(user) //序列化
	if err != nil {
		return
	}

	//入库
	_, err = conn.Do("Hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}
	return
}
