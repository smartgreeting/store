package form

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     int    `json:"code"`
}
