/*
 * @Author: lihuan
 * @Date: 2021-10-14 13:48:44
 * @LastEditTime: 2021-10-15 17:32:22
 * @Email: 17719495105@163.com
 */
package middleware

import (
	"log"
	"net/http"
	"store/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//打印错误堆栈信息
				log.Printf("panic: %v\n", r)
				// debug.PrintStack()
				//封装通用json返回
				getErrorReponse(ctx, r)
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				ctx.Abort()
			}
		}()
		//加载完 defer recover，继续后续接口调用
		ctx.Next()
	}
}

func getErrorReponse(ctx *gin.Context, r interface{}) {

	switch v := r.(type) {

	case int:
		getJson(ctx, http.StatusOK, v, utils.GetMsg(v))

	case validator.ValidationErrors:
		getJson(ctx, http.StatusOK, utils.PARAMS_PARSE_ERROR, utils.ErrorTranslate(r.(validator.ValidationErrors)))

	case string:

		getJson(ctx, http.StatusOK, utils.PARAMS_PARSE_ERROR, r.(string))

	case error:
		str := v.Error()
		code := http.StatusInternalServerError
		if strings.Contains(str, "code = Unknown desc =") {
			string_slice := strings.Split(str, "code = Unknown desc = ")
			str = string_slice[1]
			code = http.StatusOK
		}
		getJson(ctx, code, 1000, str)

	default:
		getJson(ctx, http.StatusInternalServerError, 5000, "服务异常")
	}
}

func getJson(ctx *gin.Context, httpCode, code int, msg string) {
	ctx.JSON(httpCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}
