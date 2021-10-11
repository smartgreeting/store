/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-11 16:35:56
 * @Email: 17719495105@163.com
 */
package handler

import (
	"store/api/middleware"
	"store/api/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 实例化服务
	userService := service.NewUserService()

	v1Group := r.Group("v1")
	v1Group.POST("/user/register", userService.Register)

	// 使用jtw中间件
	v1Group.Use(middleware.JWT())

	return r
}
