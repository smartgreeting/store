/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-25 13:48:08
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"
	"errors"
	"fmt"
	"store/models"
	"store/rpc-user/apiuser"
	"store/rpc-user/internal/dao"
	"store/rpc-user/internal/svc"

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
func (l *RegisterLogic) Register(in *apiuser.RegisterReq) (*apiuser.UserReply, error) {

	key := fmt.Sprintf("GetCaptuha%s", in.Phone)
	code, _ := l.svcCtx.RedCli.Get(key).Result()
	if code != in.Code {
		return nil, errors.New("验证不正确")
	}

	request := models.User{
		Phone:    in.Phone,
		Password: in.Password,
	}

	_, err := l.userDao.Create(&request)
	if err != nil {
		return nil, err
	}

	return &apiuser.UserReply{}, nil
}
