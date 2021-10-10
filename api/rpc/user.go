package rpc

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/zrpc"
	"store/rpc-user/apiuserclient"
	"store/models/in"
)

type UserRpcInterface interface {
	Register(ctx context.Context, in *in.UserRegisterReq) (interface{}, error)
}

type userRpc struct {
}

func NewUserRpc() UserRpcInterface {
	return &userRpc{}
}

func (r *userRpc) Register(ctx context.Context, in *in.UserRegisterReq) (interface{}, error) {
	rpcClientConf := zrpc.NewEtcdClientConf([]string{"127.0.0.1:2379"}, "apiuser.rpc", "", "")
	apiUser := apiuserclient.NewApiUser(zrpc.MustNewClient(rpcClientConf))
	res, err := apiUser.Register(ctx, &apiuserclient.RegisterReq{
		Phone:    "",
		Password: "",
		Code:     "",
	})
	fmt.Println("res: ", res)
	fmt.Println("err: ", err)
	return nil, nil
}
