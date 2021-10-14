/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-14 17:45:07
 * @Email: 17719495105@163.com
 */
package service

import (
	"fmt"
	"net/http"
	"store/api/rpc"
	"store/models/form"
	"store/models/in"
	"store/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// 注册
func (s *UserService) Register(ctx *gin.Context) {

	var req form.User

	err := ctx.ShouldBindWith(&req, binding.JSON)
	// 参数校验
	if err != nil {
		errs := err.(validator.ValidationErrors)
		panic(utils.ErrorTranslate(errs))
	}

	// 注册
	_, err = rpc.NewUserRpc().Register(ctx, &in.UserRegisterReq{
		Phone:    req.Phone,
		Password: utils.EncodeMd5(req.Password, []byte(utils.Cfg.Md5.Secret)),
		Code:     req.Code,
	})
	// 注册失败
	if err != nil {
		panic(utils.UserCreateErr)
	}

	utils.Response(ctx, http.StatusOK, utils.SUCCESS, nil)
}

// 登录
func (u *UserService) Login(ctx *gin.Context) {
	var req form.User

	err := ctx.ShouldBindWith(&req, binding.JSON)
	// 参数校验
	if err != nil {
		errs := err.(validator.ValidationErrors)
		panic(utils.ErrorTranslate(errs))
	}
	// 登录
	user, err := rpc.NewUserRpc().Login(ctx, &in.UserLoginReq{
		Phone:    req.Phone,
		Password: utils.EncodeMd5(req.Password, []byte(utils.Cfg.Md5.Secret)),
		Code:     req.Code,
	})
	// 登录失败
	if err != nil {
		fmt.Println(err, 111)
		panic(utils.LOGIN_ERROR)

	}
	token, err := utils.GenerateToken(int(user.ID), req.Phone, utils.EncodeMd5(req.Password, []byte(utils.Cfg.Md5.Secret)), []byte(utils.Cfg.Token.Secret), utils.Cfg.Token.ExpireTime)
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, utils.GENERATE_TOKEN_ERROR, nil)
		fmt.Println(err)
		return
	}

	utils.Response(ctx, http.StatusOK, utils.SUCCESS, gin.H{
		"id":    user.ID,
		"token": token,
	})

}
