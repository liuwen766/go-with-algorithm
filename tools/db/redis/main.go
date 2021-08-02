package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	initClient()
}

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	res, err := rdb.Ping().Result()
	fmt.Println(res)
	if err != nil {
		return err
	}
	return nil
}
