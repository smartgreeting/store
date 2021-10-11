/*
 * @Author: lihuan
 * @Date: 2021-10-11 16:10:05
 * @LastEditTime: 2021-10-11 18:00:01
 * @Email: 17719495105@163.com
 */
package middleware

import (
	"net/http"
	"store/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var code int

		code = utils.SUCCESS

		Authorization := ctx.GetHeader("Authorization")

		token := strings.Split(Authorization, " ")

		if Authorization == "" {
			code = utils.INVALID_TOKEN_ERROR
		} else {
			_, err := utils.ParseToken(token[1], []byte(utils.Cfg.Token.Secret))
			if err != nil {

				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = utils.TOKEN_TIMEOUT_ERROR
				default:
					code = utils.TOKEN_FAIL_ERROR
				}
			}
		}

		if code != utils.SUCCESS {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": utils.UNAUTHORIZED_ERROR,
				"msg":  utils.Response(utils.UNAUTHORIZED_ERROR),
				"data": nil,
			})

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
