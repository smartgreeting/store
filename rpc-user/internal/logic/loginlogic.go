/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-13 14:12:54
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"
	"errors"

	"store/models"
	"store/rpc-user/apiuser"
	"store/rpc-user/internal/dao"
	"store/rpc-user/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userDao *dao.UserDao
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		userDao: dao.NewUserDao(ctx, svcCtx.Db),
	}
}

//   登录接口
func (l *LoginLogic) Login(in *apiuser.LoginReq) (*apiuser.LoginRes, error) {
	// 短信验证码登录
	// if len(in.Password) == 0 {
	// 	//	:todo
	// }

	// 账号密码登录
	user, err := l.checkPassword(in.Password, in.Phone)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, errors.New("密码错误")
	}

	return &apiuser.LoginRes{
		Id: user.ID,
	}, nil
}

func (l *LoginLogic) checkPassword(reqPassword, phone string) (*models.User, error) {
	user, err := l.userDao.FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	if user.Password == reqPassword {
		return user, nil
	}
	return nil, errors.New("")

}
