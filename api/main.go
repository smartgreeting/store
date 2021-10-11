package main

import (
	"fmt"

	"store/api/handler"
	"store/utils"
)

func main() {
	cfg, err := utils.GetConf("./conf/conf.yaml")
	if err != nil {
		panic(err.Error())
	}
	r := handler.SetupRouter()
	r.Run(fmt.Sprintf("%s:%d", cfg.Application.Address, cfg.Application.Port))
}
