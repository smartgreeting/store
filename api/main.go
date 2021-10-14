/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-14 13:57:46
 * @Email: 17719495105@163.com
 */
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
