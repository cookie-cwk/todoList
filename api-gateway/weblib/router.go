package weblib

import (
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)
func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"code": 200,
				"msg": "success",
			})
		})
		v1.POST("/user/register",handlers.UserRegister)
		v1.POST("/user/login",handlers.UserLogin)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("tasks", handlers.GetTaskList)
			authed.POST("task", handlers.CreateTask)
			authed.GET("task/:id", handlers.GetTaskDetail)
			authed.PUT("task/:id", handlers.UpdateTask)
			authed.DELETE("task/:id", handlers.DeleteTask)
		}
	}
	return ginRouter
}
