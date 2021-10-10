// Code generated by goctl. DO NOT EDIT!
// Source: api-user.proto

package server

import (
	"context"

	"store/rpc-user/apiuser"
	"store/rpc-user/internal/logic"
	"store/rpc-user/internal/svc"
)

type ApiUserServer struct {
	svcCtx *svc.ServiceContext
}

func NewApiUserServer(svcCtx *svc.ServiceContext) *ApiUserServer {
	return &ApiUserServer{
		svcCtx: svcCtx,
	}
}

//   注册接口
func (s *ApiUserServer) Register(ctx context.Context, in *apiuser.RegisterReq) (*apiuser.NullRes, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

//   登录接口
func (s *ApiUserServer) Login(ctx context.Context, in *apiuser.LoginReq) (*apiuser.LoginRes, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

//   获取用户信息
func (s *ApiUserServer) GetUser(ctx context.Context, in *apiuser.GetUserReq) (*apiuser.GetUserRes, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

//   更新用户信息
func (s *ApiUserServer) UpdateUser(ctx context.Context, in *apiuser.UpdateUserReq) (*apiuser.UpdateUserRes, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}
