package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	addr := "127.0.0.1:6379"
	redis, _ := NewRds(addr, "", 0, 10)

	err := redis.StringSet(time.Second*time.Duration(100), "name", "xy")
	if err != nil {
		fmt.Println("error", err)
	}

	value := redis.Get("name")
	fmt.Println(value)
}

type clientRedis struct {
	*redis.Client
}

/*
   获取一个*redis.Client
*/
func NewRds(addr, password string, DB, poolSize int) (client *clientRedis, err error) {
	client = &clientRedis{redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       DB,
		PoolSize: poolSize,
	})}

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pong)
	return
}

/*
   string set操作  对照命令 set key value
*/
func (cr *clientRedis) StringSet(expire time.Duration, key string, value string) (err error) {
	err = cr.Client.Append(key, value).Err()
	if err != nil {
		// todo error info
		return
	}
	cr.Expire(key, expire)
	return
}

/*
   string get操作  对照命令 get key
*/
func (cr *clientRedis) StringGet(key string) (res string, err error) {
	res, err = cr.Get(key).Result()
	if err != nil {
		//todo error info
		return
	}
	return
}

/*
   list lpush操作 对照命令  lpush key value
*/
func (cr *clientRedis) ListLPush(expire time.Duration, key string, value ...string) (err error) {
	err = cr.LPush(key, value).Err()
	if err != nil {
		// todo error info
		return
	}
	cr.Expire(key, expire)
	return
}

/*
   list rpush操作 对照命令  rpush key value
*/
func (cr *clientRedis) ListRPush(expire time.Duration, key string, value ...string) (err error) {
	err = cr.RPush(key, value).Err()
	if err != nil {
		// todo error info
		return
	}
	cr.Expire(key, expire)
	return
}

/*
   list lpop操作  对照命令  lpop key
*/
func (cr *clientRedis) ListLPop(key string) (res string, err error) {
	res, err = cr.LPop(key).Result()
	if err != nil {
		//todo error info
		return
	}
	return
}

/*
   list rpop操作  对照命令  rpop key
*/
func (cr *clientRedis) ListRPop(key string) (res string, err error) {
	res, err = cr.RPop(key).Result()
	if err != nil {
		//todo error info
		return
	}
	return
}

/*
   set  add操作  对照命令   sadd key value
*/
func (cr *clientRedis) SetAdd(expire time.Duration, key string, value ...string) (err error) {
	err = cr.SAdd(key, value).Err()
	if err != nil {
		// todo error info
		return
	}
	cr.Expire(key, expire)
	return
}

/*
   set members操作 对照命令  smembers key
*/
func (cr *clientRedis) SetMembers(key string) (res []string, err error) {
	res, err = cr.SMembers(key).Result()
	if err != nil {
		//todo error info
		return
	}
	return
}
