package handlers

import (
	"api-gateway/pkg/utils"
	service "api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(c.Bind(&userReq))

	userService := c.Keys["userService"].(service.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	c.JSON(http.StatusOK, gin.H{"data": userResp})
}


func UserLogin(c *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(c.Bind(&userReq))

	userService := c.Keys["userService"].(service.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))
	c.JSON(http.StatusOK, gin.H{
		"code": userResp.Code,
		"data": gin.H{
			"user": userResp,
			"token": token,
		},
	})
}