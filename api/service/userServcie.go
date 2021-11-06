/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-26 15:41:04
 * @Email: 17719495105@163.com
 */
package service

import (
	"fmt"
	"net/http"
	"regexp"
	"store/api/rpc"
	"store/models"
	"store/models/form"
	"store/rpc-user/apiuser"
	"store/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

		panic(err)

	}

	// 注册
	_, err = rpc.NewUserRpc().Register(ctx, &apiuser.RegisterReq{
		Phone:    req.Phone,
		Password: utils.EncodeMd5(req.Password, []byte(utils.Cfg.Md5.Secret)),
		Code:     req.Code,
	})
	// 注册失败
	if err != nil {
		// panic(utils.UserCreateErr)
		panic(err)
		
	}

	utils.Response(ctx, http.StatusOK, utils.SUCCESS, nil)
}

// 登录
func (u *UserService) Login(ctx *gin.Context) {
	var req form.User

	err := ctx.ShouldBindWith(&req, binding.JSON)
	// 参数校验
	if err != nil {
		panic(err)
	}
	// 登录
	user, err := rpc.NewUserRpc().Login(ctx, &apiuser.LoginReq{
		Phone:    req.Phone,
		Password: utils.EncodeMd5(req.Password, []byte(utils.Cfg.Md5.Secret)),
		Code:     req.Code,
	})
	// 登录失败
	if err != nil {

		panic(err)
		// panic(utils.LOGIN_ERROR)

	}
	token, err := utils.GenerateToken(int(user.Id), req.Phone, utils.EncodeMd5(req.Password, []byte(utils.Cfg.Md5.Secret)), []byte(utils.Cfg.Token.Secret), utils.Cfg.Token.ExpireTime)
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, utils.GENERATE_TOKEN_ERROR, nil)
		fmt.Println(err)
		return
	}

	utils.Response(ctx, http.StatusOK, utils.SUCCESS, gin.H{
		"id":    user.Id,
		"token": token,
	})

}

func (u *UserService) GetCaptuha(ctx *gin.Context) {
	phone := ctx.Query("phone")
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	if !rgx.MatchString(phone) {
		panic(utils.ErrorPhoneNotExit)
	}
	res, err := rpc.NewUserRpc().GetCaptuha(ctx, &apiuser.GetCaptuhaReq{
		Phone: phone,
	})
	if err != nil {
		panic(err)
	}
	utils.Response(ctx, http.StatusOK, utils.SUCCESS, gin.H{
		"code": res.Code,
	})
}

func (u *UserService) GetUserInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))

	res, err := rpc.NewUserRpc().GetUserInfo(ctx, &apiuser.GetUserReq{Id: uint64(id)})
	if err != nil {
		panic(err)
	}
	
	utils.Response(ctx, http.StatusOK, utils.SUCCESS, gin.H{
		
		"userInfo": models.RpcToUser(res),
	})
}
func (u *UserService) UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var req form.UserInfo

	err := ctx.ShouldBindWith(&req, binding.JSON)
	// 参数校验
	if err != nil {
		panic(err)
	}

	res, err := rpc.NewUserRpc().UpdateUser(ctx, &apiuser.UpdateUserReq{
		User: &apiuser.User{
			Id:       uint64(id),
			Username: req.Username,
			Avatar:   req.Avatar,
			Sex:      req.Sex,
			Phone:    req.Phone,
			Email:    req.Email,
			Address:  req.Address,
			Hobby:    req.Hobby,
		},
	})
	if err != nil {
		panic(err)
	}
	utils.Response(ctx, http.StatusOK, utils.SUCCESS, gin.H{
		"userInfo": models.RpcToUser(res),
	})
}
