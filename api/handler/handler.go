/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-25 16:12:58
 * @Email: 17719495105@163.com
 */
package handler

import (
	"store/api/middleware"
	"store/api/service"
	"store/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 注册参数校验器
	utils.Validator()

	r.Use(middleware.Recover())

	// 实例化服务
	userService := service.NewUserService()

	v1Group := r.Group("v1")
	v1Group.POST("/user/register", userService.Register)
	v1Group.POST("/user/login", userService.Login)
	v1Group.GET("/user/captuha", userService.GetCaptuha)

	// 使用jtw中间件
	v1Group.Use(middleware.JWT())

	{
		v1Group.GET("/user/userInfo", userService.GetUserInfo)
		v1Group.PUT("/user/userInfo/:id", userService.UpdateUser)

	}

	return r
}
