/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-25 11:31:13
 * @Email: 17719495105@163.com
 */
package utils

import (
	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = iota + 2000
	UserCreateErr
	DateBaseErr
	INVALID_TOKEN_ERROR
	TOKEN_TIMEOUT_ERROR
	TOKEN_FAIL_ERROR
	UNAUTHORIZED_ERROR
	PARAMS_PARSE_ERROR
	LOGIN_ERROR
	GENERATE_TOKEN_ERROR
	ErrorPhoneNotExit
)

var resMap = map[int]string{
	SUCCESS:              "成功",
	UserCreateErr:        "用户创建失败",
	DateBaseErr:          "数据库错误",
	INVALID_TOKEN_ERROR:  "无效的 Token",
	TOKEN_TIMEOUT_ERROR:  "token过期",
	TOKEN_FAIL_ERROR:     "解析token失败",
	UNAUTHORIZED_ERROR:   "没有权限",
	PARAMS_PARSE_ERROR:   "请求参数解析错误",
	LOGIN_ERROR:          "登录失败",
	GENERATE_TOKEN_ERROR: "生成token失败",
	ErrorPhoneNotExit:    "手机号码不正确！",
}

func Response(ctx *gin.Context, httpCode, code int, data interface{}) {
	
	ctx.JSON(httpCode, gin.H{
		"code": code,
		"msg":  GetMsg(code),
		"data": data,
	})

}

func GetMsg(code int) string {

	str, ok := resMap[code]

	if ok {
		return str
	}
	return "未知错误"
}
