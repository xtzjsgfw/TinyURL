package main

import (
	"TinyURL/extend/conf"
	"TinyURL/extend/redis"
	"TinyURL/models"
	"TinyURL/routes"
)

func main() {
	// 初始化配置
	conf.Setup()

	// 初始化数据库
	models.Setup()

	// 初始化redis
	redis.Setup()

	// 初始化路由
	routes.Setup()
}

