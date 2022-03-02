package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//配置跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerkeys []string
		for k := range c.Request.Header {
			headerkeys = append(headerkeys, k)
		}
		headerStr := strings.Join(headerkeys, ",")

		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin", "access-control-allow-header, %s",headerStr)
		} else {
			headerStr = fmt.Sprintln("access-control-allow-origin", "access-control-allow-header")
		}
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "*")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "false")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}