package middleware

import (
	"net/http"
	"task4/authhandle"

	"github.com/gin-gonic/gin"
)

func LoginMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 这里的Token放在请求头里
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 10501,
				"msg":  "您还未登录,请先登录!",
			})
			c.Abort()
			return
		}
		mc, err := authhandle.ParseLoginToken(authHeader)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 10502,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
