// Code generated by goctl. DO NOT EDIT!
// Source: api-user.proto

package apiuserclient

import (
	"context"

	"store/rpc-user/apiuser"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	User          = apiuser.User
	UpdateUserReq = apiuser.UpdateUserReq
	UserReply     = apiuser.UserReply
	RegisterReq   = apiuser.RegisterReq
	LoginReq      = apiuser.LoginReq
	GetCaptuhaReq = apiuser.GetCaptuhaReq
	GetUserReq    = apiuser.GetUserReq

	ApiUser interface {
		//   注册接口
		Register(ctx context.Context, in *RegisterReq) (*UserReply, error)
		//   登录接口
		Login(ctx context.Context, in *LoginReq) (*UserReply, error)
		//   获取用户信息
		GetUser(ctx context.Context, in *GetUserReq) (*UserReply, error)
		//   更新用户信息
		UpdateUser(ctx context.Context, in *UpdateUserReq) (*UserReply, error)
		//  获取验证码
		GetCaptuha(ctx context.Context, in *GetCaptuhaReq) (*UserReply, error)
	}

	defaultApiUser struct {
		cli zrpc.Client
	}
)

func NewApiUser(cli zrpc.Client) ApiUser {
	return &defaultApiUser{
		cli: cli,
	}
}

//   注册接口
func (m *defaultApiUser) Register(ctx context.Context, in *RegisterReq) (*UserReply, error) {
	client := apiuser.NewApiUserClient(m.cli.Conn())
	return client.Register(ctx, in)
}

//   登录接口
func (m *defaultApiUser) Login(ctx context.Context, in *LoginReq) (*UserReply, error) {
	client := apiuser.NewApiUserClient(m.cli.Conn())
	return client.Login(ctx, in)
}

//   获取用户信息
func (m *defaultApiUser) GetUser(ctx context.Context, in *GetUserReq) (*UserReply, error) {
	client := apiuser.NewApiUserClient(m.cli.Conn())
	return client.GetUser(ctx, in)
}

//   更新用户信息
func (m *defaultApiUser) UpdateUser(ctx context.Context, in *UpdateUserReq) (*UserReply, error) {
	client := apiuser.NewApiUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in)
}

//  获取验证码
func (m *defaultApiUser) GetCaptuha(ctx context.Context, in *GetCaptuhaReq) (*UserReply, error) {
	client := apiuser.NewApiUserClient(m.cli.Conn())
	return client.GetCaptuha(ctx, in)
}
