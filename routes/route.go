package routes

import (
	v1 "TinyURL/controller/v1"
	"TinyURL/middleware"
	"github.com/gin-gonic/gin"
)

func Setup() {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	var urlController v1.URLController
	router := engine.Group("api")
	{
		var authController v1.AuthController

		router.POST("/auth/sendcode", authController.SendCode)
		router.POST("/auth/register", authController.Register)
		router.POST("/auth/login", authController.Login)
		router.POST("/url/create", urlController.GenerateTinyUrlByIncreaseNum)

		router.Use(middleware.JWTAuth())
		{
			// 写一些需要权限的接口
		}
	}
	// 短链接重定向接口
	engine.Use(middleware.IgnoreURL())
	engine.GET("/*tinyUrl", urlController.RequestByTinyUrl)
	engine.Run()
}

