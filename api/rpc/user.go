package rpc

import (
	"context"
	"fmt"
	"store/models/in"
	"store/rpc-user/apiuserclient"
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
	apiUser := apiuserclient.NewApiUser(RpcClient)
	res, err := apiUser.Register(ctx, &apiuserclient.RegisterReq{
		Phone:    "",
		Password: "",
		Code:     "",
	})
	fmt.Println("res: ", res)
	fmt.Println("err: ", err)
	return nil, nil
}
