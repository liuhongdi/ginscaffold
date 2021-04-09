package middleware

import (
	"fmt"
	"ginscaffold/global"
	"ginscaffold/pkg/jwt"
	"ginscaffold/pkg/result"
	"github.com/gin-gonic/gin"
	"strings"
)

//基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在URI的Authorization中，并使用Bearer开头

		authHeader := c.Request.Header.Get("Authorization")
		/*
		fmt.Println("begin JWTAuthMiddleware")
		auth:= c.Query("auth")
		fmt.Println("auth")
		fmt.Println(auth)
		authHeader := auth
		*/
		if authHeader == "" {
			result := result.NewResult(c)
			result.Error(global.ErrUserNoHeader)
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			result := result.NewResult(c)
			result.Error(global.ErrUserNoAuth)
			return
		}
		fmt.Println("tokenString")
		fmt.Println(parts[1])
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			result := result.NewResult(c)
			result.Error(global.ErrUserToken)
			return
		}
		fmt.Println("mc:")
		fmt.Println(mc.Username)
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Next()
		// 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
