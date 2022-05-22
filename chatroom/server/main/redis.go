package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool //创建redis连接池

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle:     maxIdle,     //最初的连接数量
		MaxActive:   maxActive,   //最大连接数量
		IdleTimeout: idleTimeout, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", address)
		},
	}
}
