/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-25 16:00:58
 * @Email: 17719495105@163.com
 */
package rpc

import (
	"context"
	"store/models/in"
	"store/models/reply"
	"store/rpc-user/apiuser"
	"store/rpc-user/apiuserclient"

	"github.com/tal-tech/go-zero/zrpc"
)

type UserRpcInterface interface {
	Register(ctx context.Context, in *in.UserRegisterReq) (*reply.UserReply, error)
	Login(ctx context.Context, in *in.UserLoginReq) (*reply.UserReply, error)
	GetCaptuha(ctx context.Context, in *apiuser.GetCaptuhaReq) (string, error)
	GetUserInfo(ctx context.Context, in *apiuser.GetUserReq) (*reply.UserReply, error)
	UpdateUser(ctx context.Context, in *apiuser.UpdateUserReq) (*reply.UserReply, error)
}

type userRpc struct {
}

var apiuserRpc apiuserclient.ApiUser

func init() {
	rpcClientConf := zrpc.NewEtcdClientConf([]string{"127.0.0.1:2379"}, "apiuser.rpc", "", "")
	client := zrpc.MustNewClient(rpcClientConf)
	apiuserRpc = apiuserclient.NewApiUser(client)
}

func NewUserRpc() UserRpcInterface {
	return &userRpc{}
}

func (r *userRpc) Register(ctx context.Context, in *in.UserRegisterReq) (*reply.UserReply, error) {

	_, err := apiuserRpc.Register(ctx, &apiuserclient.RegisterReq{
		Phone:    in.Phone,
		Password: in.Password,
		Code:     in.Code,
	})
	if err != nil {
		return nil, err
	}

	return &reply.UserReply{}, nil
}

func (r *userRpc) Login(ctx context.Context, in *in.UserLoginReq) (*reply.UserReply, error) {

	user, err := apiuserRpc.Login(ctx, &apiuserclient.LoginReq{
		Phone:    in.Phone,
		Password: in.Password,
		Code:     in.Code,
	})

	if err != nil {
		return nil, err
	}

	return &reply.UserReply{
		ID: user.Id,
	}, nil

}

func (r *userRpc) GetCaptuha(ctx context.Context, in *apiuser.GetCaptuhaReq) (string, error) {
	res, err := apiuserRpc.GetCaptuha(ctx, &apiuser.GetCaptuhaReq{
		Phone: in.Phone,
	})
	if err != nil {
		return "", err
	}
	return res.Code, nil
}

func (r *userRpc) GetUserInfo(ctx context.Context, in *apiuser.GetUserReq) (*reply.UserReply, error) {
	res, err := apiuserRpc.GetUser(ctx, &apiuser.GetUserReq{
		Id: in.Id,
	})
	if err != nil {
		return nil, err
	}
	return &reply.UserReply{
		ID:       res.Id,
		Username: res.Username,
		Avatar:   res.Avatar,
		Sex:      int32(res.Sex),
		Phone:    res.Phone,
		Email:    res.Email,
		Address:  res.Address,
		Hobby:    res.Hobby,
	}, nil
}
func (r *userRpc) UpdateUser(ctx context.Context, in *apiuser.UpdateUserReq) (*reply.UserReply, error) {
	res, err := apiuserRpc.UpdateUser(ctx, &apiuser.UpdateUserReq{
		User: in.User,
	})
	if err != nil {
		return nil, err
	}
	return &reply.UserReply{
		ID:       res.Id,
		Username: res.Username,
		Avatar:   res.Avatar,
		Sex:      int32(res.Sex),
		Phone:    res.Phone,
		Email:    res.Email,
		Address:  res.Address,
		Hobby:    res.Hobby,
	}, nil
}
