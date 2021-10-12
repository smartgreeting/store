package main

import (
	"fmt"

	"store/api/handler"
	"store/utils"

	"github.com/tal-tech/go-zero/zrpc"
)

func main() {
	cfg, err := utils.GetConf("./conf/conf.yaml")
	rpcClientConf := zrpc.NewEtcdClientConf([]string{"127.0.0.1:2379"}, "apiuser.rpc", "", "")
	client := zrpc.MustNewClient(rpcClientConf)

	if err != nil {
		panic(err.Error())
	}
	r := handler.SetupRouter(client)
	r.Run(fmt.Sprintf("%s:%d", cfg.Application.Address, cfg.Application.Port))
}
