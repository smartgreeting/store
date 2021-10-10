package svc

import (
	"fmt"
	"github.com/go-redis/redis"
	"store/rpc-user/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
	RedCli *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 改成配置文件
	db, err := gorm.Open(mysql.Open(c.Mysql.Dns), &gorm.Config{})

	redisDb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port), // redis地址
		Password: "",                                               // redis密码，没有则留空
		DB:       c.Redis.DB,                                       // 默认数据库，默认是0
	})

	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		Db:     db,
		RedCli: redisDb,
	}
}
