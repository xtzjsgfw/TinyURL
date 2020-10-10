package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func Setup() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		//DB: 0,
	})
	pone, err := Rdb.Ping().Result()
	if err != nil {
		fmt.Println("连接redis失败")
	}
	fmt.Println("连接成功", pone)
}