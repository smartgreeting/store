/*
 * @Author: lihuan
 * @Date: 2021-10-11 13:10:44
 * @LastEditTime: 2021-10-25 15:21:20
 * @Email: 17719495105@163.com
 */
package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserName string `json:"username"`
	ID       int    `json:"id"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(id int, username, secret_password string, secret []byte, exp int) (string, error) {

	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(exp) * time.Hour)

	claims := Claims{
		username,
		id,
		secret_password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "smartgreeting",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(secret)

	return token, err
}

func ParseToken(token string, secret []byte) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
