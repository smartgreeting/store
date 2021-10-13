/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-13 14:16:02
 * @Email: 17719495105@163.com
 */
package rpc

import (
	"context"
	"store/models/in"
	"store/rpc-user/apiuserclient"

	"github.com/tal-tech/go-zero/zrpc"
)

type UserRpcInterface interface {
	Register(ctx context.Context, in *in.UserRegisterReq) (interface{}, error)
	Login(ctx context.Context, in *in.UserLoginReq) (uint64, error)
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

func (r *userRpc) Register(ctx context.Context, in *in.UserRegisterReq) (interface{}, error) {

	_, err := apiuserRpc.Register(ctx, &apiuserclient.RegisterReq{
		Phone:    in.Phone,
		Password: in.Password,
		Code:     in.Code,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *userRpc) Login(ctx context.Context, in *in.UserLoginReq) (uint64, error) {

	res, err := apiuserRpc.Login(ctx, &apiuserclient.LoginReq{
		Phone:    in.Phone,
		Password: in.Password,
		Code:     in.Code,
	})

	if err != nil {
		return 0, err
	}

	return res.Id, nil

}
