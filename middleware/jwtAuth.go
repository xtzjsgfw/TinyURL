package middleware

import (
	"TinyURL/extend/code"
	"TinyURL/extend/jwt"
	"TinyURL/extend/redis"
	"TinyURL/extend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// JWTAuth Token 认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 Authorization token 值
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			fmt.Println("未获取到Token值")
			utils.ResponseFormat(c, code.TokenNotFound, nil)
			c.Abort()
			return
		}

		// 获取到 Token, 解析token信息
		token = token[7:]
		fmt.Println(token)
		jwtInstance := jwt.NewJWT()
		claims, err := jwtInstance.ParseToken(token)
		if err != nil {
			// 未能正常解析 Token，则报：token认证失败
			fmt.Println("无法正常解析Token")
			utils.ResponseFormat(c, code.TokenInvalid, nil)
			c.Abort()
			return
		}
		fmt.Println(claims)

		// 获取缓存中的Token信息
		tokenCache, err := redis.Rdb.Get("TOKEN:" + claims.Phone).Result()
		if err != nil {
			fmt.Printf("jwt auth redis get: %v", err.Error())
			fmt.Println("Redis中没有Token信息")
			utils.ResponseFormat(c, code.ServiceInsideError, nil)
			c.Abort()
			return
		}

		// 用户注销或token失效
		if tokenCache != token {
			fmt.Println("Token无效")
			utils.ResponseFormat(c, code.TokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
