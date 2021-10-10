package utils

const (
	Success = iota + 2000
	UserCreateErr
	DateBaseErr
)

var resMap = map[int]string{
	UserCreateErr: "用户创建失败",
	DateBaseErr:   "数据库错误",
}

func Response(code int) string {
	if _, ok := resMap[code]; ok {
		return "未知错误"
	}

	return resMap[code]
}
