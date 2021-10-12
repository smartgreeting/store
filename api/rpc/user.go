package rpc

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/zrpc"
	"store/models/in"
	"store/rpc-user/apiuserclient"
)

type UserRpcInterface interface {
	Register(ctx context.Context, in *in.UserRegisterReq) (interface{}, error)
}

type userRpc struct {
	client zrpc.Client
}

func NewUserRpc(client zrpc.Client) UserRpcInterface {
	return &userRpc{client}
}

func (r *userRpc) Register(ctx context.Context, in *in.UserRegisterReq) (interface{}, error) {
	apiUser := apiuserclient.NewApiUser(r.client)
	res, err := apiUser.Register(ctx, &apiuserclient.RegisterReq{
		Phone:    "",
		Password: "",
		Code:     "",
	})
	fmt.Println("res: ", res)
	fmt.Println("err: ", err)
	return nil, nil
}
