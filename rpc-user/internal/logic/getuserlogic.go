/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-25 15:43:46
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

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userDao *dao.UserDao
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		userDao: dao.NewUserDao(ctx, svcCtx.Db),
	}
}

//   获取用户信息
func (l *GetUserLogic) GetUser(in *apiuser.GetUserReq) (*apiuser.UserReply, error) {
	userInfo, err := l.userDao.FindUserInfoById(in.Id)
	switch err {
	case nil:

		return &apiuser.UserReply{
			Id:       userInfo.ID,
			Username: userInfo.Username,
			Avatar:   userInfo.Avatar,
			Sex:      int32(userInfo.Sex),
			Phone:    userInfo.Phone,
			Email:    userInfo.Email,
			Address:  userInfo.Address,
			Hobby:    userInfo.Hobby,
		}, nil
	case gorm.ErrRecordNotFound:
		return nil, errors.New("未匹配到记录")
	default:
		return nil, err
	}

}
