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
func (l *GetUserLogic) GetUser(in *apiuser.GetUserReq) (*apiuser.GetUserRes, error) {
	// todo: add your logic here and delete this line

	return &apiuser.GetUserRes{}, nil
}
