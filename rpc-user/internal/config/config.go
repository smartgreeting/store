package config

import (
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql Mysql
	Redis Redis
	UserAuth  Auth
}

type Mysql struct {
	Dns string
}

type Redis struct {
	Host string
	Port int
	DB   int
}

type Auth struct {
	Md5 string
}
