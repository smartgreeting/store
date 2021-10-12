package service

import (
	"store/api/rpc"
	"store/models/form"
	"store/models/in"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register(ctx *gin.Context) {
	var req form.User
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    "",
			"message": "hello",
			"data":    "ok",
		})

		return
	}

	_, err = rpc.NewUserRpc().Register(ctx, &in.UserRegisterReq{
		Phone:    "123",
		Password: "123",
		Code:     "123",
	})

	ctx.JSON(200, gin.H{
		"message": "hello",
	})
}
