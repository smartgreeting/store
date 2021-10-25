/*
 * @Author: lihuan
 * @Date: 2021-10-25 10:42:35
 * @LastEditTime: 2021-10-25 10:59:50
 * @Email: 17719495105@163.com
 */
package logic

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"store/rpc-user/apiuser"
	"store/rpc-user/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCaptuhaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCaptuhaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptuhaLogic {
	return &GetCaptuhaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取验证码
func (l *GetCaptuhaLogic) GetCaptuha(in *apiuser.GetCaptuhaReq) (*apiuser.UserReply, error) {
	key := fmt.Sprintf("GetCaptuha%s", in.Phone)
	code := genCode(6)
	err := l.svcCtx.RedCli.Set(key, code, 5*time.Minute).Err()
	if err != nil {
		return nil, errors.New("生成验证码失败")
	}

	return &apiuser.UserReply{
		Code: code,
	}, nil
}

func genCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()

}
