package middleware

import (
	"api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 401
		} else {
			if strings.HasPrefix(token, "Bearer ") {
				 token = token[7:]
			}
			_, err := utils.ParseToken(token)
			if  err != nil {
				code = 401
			}
		}
		if code != 200 {
			c.JSON(code, gin.H{
				"code": code,
				"msg": "鉴权失败",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
