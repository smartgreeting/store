package utils

const (
	SUCCESS = iota + 2000
	UserCreateErr
	DateBaseErr
	INVALID_TOKEN_ERROR
	TOKEN_TIMEOUT_ERROR
	TOKEN_FAIL_ERROR
	UNAUTHORIZED_ERROR
)

var resMap = map[int]string{
	UserCreateErr:       "用户创建失败",
	DateBaseErr:         "数据库错误",
	INVALID_TOKEN_ERROR: "无效的 Token",
	TOKEN_TIMEOUT_ERROR: "token过期",
	TOKEN_FAIL_ERROR:    "解析token失败",
	UNAUTHORIZED_ERROR:  "没有权限",
}

func Response(code int) string {
	if _, ok := resMap[code]; ok {
		return "未知错误"
	}

	return resMap[code]
}
