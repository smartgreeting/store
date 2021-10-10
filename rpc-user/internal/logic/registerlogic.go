package logic

import (
	"context"
	"store/rpc-user/apiuser"
	"store/rpc-user/internal/dao"
	"store/rpc-user/internal/svc"
	"store/rpc-user/utils"
	"store/models"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userDao *dao.UserDao
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		userDao: dao.NewUserDao(ctx, svcCtx.Db),
	}
}

//   注册接口
func (l *RegisterLogic) Register(in *apiuser.RegisterReq) (*apiuser.NullRes, error) {
	request := models.User{
		Phone:    in.Phone,
		Password: utils.EncodeMd5([]byte(in.Password), []byte(l.svcCtx.Config.UserAuth.Md5)),
	}

	_, err := l.userDao.Create(&request)
	if err != nil {
		return nil, err
	}

	return &apiuser.NullRes{}, nil
}
