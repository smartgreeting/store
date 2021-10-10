package logic

import (
	"context"
	"store/rpc-user/apiuser"
	"store/rpc-user/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//   登录接口
func (l *LoginLogic) Login(in *apiuser.LoginReq) (*apiuser.LoginRes, error) {
	// 短信验证码登录
	if len(in.Password) == 0 {

	}

	// 账号密码登录
	return &apiuser.LoginRes{}, nil
}
