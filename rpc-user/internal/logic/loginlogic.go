package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"store/rpc-user/apiuser"
	"store/rpc-user/internal/dao"
	"store/rpc-user/internal/svc"
	"store/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userDao *dao.UserDao
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:     ctx,
		svcCtx:  svcCtx,
		Logger:  logx.WithContext(ctx),
		userDao: dao.NewUserDao(ctx, svcCtx.Db),
	}
}

//   登录接口
func (l *LoginLogic) Login(in *apiuser.LoginReq) (*apiuser.LoginRes, error) {
	// 短信验证码登录
	if len(in.Password) == 0 {
		//	:todo
	}

	// 账号密码登录
	password := utils.EncodeMd5([]byte(in.Password), []byte(l.svcCtx.Config.UserAuth.Md5))
	pass, err := l.checkPassword(password, in.Phone)
	if err != nil {
		return nil, err
	}

	if !pass {
		return nil, errors.New("密码错误")
	}

	// 生成token
	keyInfo := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()"
	info := map[string]interface{}{}
	info["user_id"] = "张三"
	dataByte, _ := json.Marshal(info)
	var dataStr = string(dataByte)

	// 使用Claim保存json
	claim := jwt.StandardClaims{Subject: dataStr, ExpiresAt: time.Now().Add(24 * time.Hour).Unix()}
	tokenInfo := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := tokenInfo.SignedString([]byte(keyInfo))
	if err != nil {
		return nil, err
	}

	// 将token字符串转换为token对象
	tokenInfo, _ = jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return keyInfo, nil
	})

	err = tokenInfo.Claims.Valid()
	if err != nil {
		println(err.Error())
	}

	finToken := tokenInfo.Claims.(jwt.MapClaims)
	// 校验下token是否过期
	succ := finToken.VerifyExpiresAt(time.Now().Unix(), true)
	fmt.Println("succ", succ)
	// 获取token中保存的用户信息
	fmt.Println(finToken["sub"])

	return &apiuser.LoginRes{}, nil
}

func (l *LoginLogic) checkPassword(reqPassword, phone string) (bool, error) {
	user, err := l.userDao.FindByPhone(phone)
	if err != nil {
		return false, err
	}

	return user.Password == reqPassword, nil
}
