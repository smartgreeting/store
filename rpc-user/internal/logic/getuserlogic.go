/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-14 11:29:35
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"
	"store/rpc-user/apiuser"
	"store/rpc-user/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//   获取用户信息
func (l *GetUserLogic) GetUser(in *apiuser.GetUserReq) (*apiuser.UserReply, error) {
	// todo: add your logic here and delete this line

	return &apiuser.UserReply{}, nil
}
