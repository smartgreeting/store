/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-14 11:29:41
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"
	"store/rpc-user/apiuser"
	"store/rpc-user/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//   更新用户信息
func (l *UpdateUserLogic) UpdateUser(in *apiuser.UpdateUserReq) (*apiuser.UserReply, error) {
	// todo: add your logic here and delete this line

	return &apiuser.UserReply{}, nil
}
