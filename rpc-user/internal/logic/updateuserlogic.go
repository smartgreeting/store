/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-25 16:16:18
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"

	"store/models"
	"store/rpc-user/apiuser"
	"store/rpc-user/internal/dao"
	"store/rpc-user/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userDao *dao.UserDao
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		userDao: dao.NewUserDao(ctx, svcCtx.Db),
	}
}

//   更新用户信息
func (l *UpdateUserLogic) UpdateUser(in *apiuser.UpdateUserReq) (*apiuser.UserReply, error) {
	req := models.User{

		ID:       in.User.Id,
		Username: in.User.Username,
		Avatar:   in.User.Avatar,
		Sex:      int(in.User.Sex),
		Phone:    in.User.Phone,
		Email:    in.User.Email,
		Address:  in.User.Address,
		Hobby:    in.User.Hobby,
	}
	userInfo, err := l.userDao.UpdateUser(&req)
	if err != nil {
		return nil, err
	}
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
}
