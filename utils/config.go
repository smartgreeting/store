/*
 * @Author: lihuan
 * @Date: 2021-10-11 11:15:50
 * @LastEditTime: 2021-10-13 09:35:29
 * @Email: 17719495105@163.com
 */
package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Application Application
	Token       Token
	Md5         Md5
}

type Token struct {
	Secret     string
	ExpireTime int
}
type Md5 struct {
	Secret string
}

type Application struct {
	Address string
	Port    int
	Mode    string
}

var Cfg *Conf = nil

func GetConf(path string) (*Conf, error) {

	// 加载文件
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 将读取的yaml文件解析为响应的 struct
	err = yaml.Unmarshal(yamlFile, &Cfg)

	return Cfg, err
}
