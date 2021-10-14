/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-14 17:53:25
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"
	"errors"

	"store/rpc-user/apiuser"
	"store/rpc-user/internal/dao"
	"store/rpc-user/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
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
func (l *LoginLogic) Login(in *apiuser.LoginReq) (*apiuser.UserReply, error) {

	user, err := l.userDao.FindByPhone(in.Phone)

	switch err {
	case nil:
		if user.Password != in.Password {
			return nil, errors.New("密码错误")
		}
		return &apiuser.UserReply{
			Id: user.ID,
		}, nil
	case gorm.ErrRecordNotFound:
		return nil, errors.New("未匹配到记录")
	default:
		return nil, err
	}

}
