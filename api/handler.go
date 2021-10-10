package main

import (
	"github.com/gin-gonic/gin"
	"store/api/service"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// 实例化服务
	userService := service.NewUserService()

	v1Group := r.Group("v1")
	v1Group.POST("/user/register", userService.Register)

	return r
}
