package rpc

import "github.com/tal-tech/go-zero/zrpc"

var RpcClient zrpc.Client

func init() {
	rpcClientConf := zrpc.NewEtcdClientConf([]string{"127.0.0.1:2379"}, "apiuser.rpc", "", "")
	RpcClient = zrpc.MustNewClient(rpcClientConf)
}
