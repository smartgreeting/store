/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-25 18:00:30
 * @Email: 17719495105@163.com
 */
package form

type User struct {
	Phone    string `json:"Phone" binding:"required,max=30,min=1" label:"手机号"`
	Password string `json:"password" binding:"required,max=60,min=6" label:"密码"`
	Code     string `json:"code" binding:"required,max=6,min=4" label:"验证码"`
}

type UserInfo struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	Sex       int32  `json:"sex"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Hobby     string `json:"hobby"`
	Deleted   int    `json:"deleted"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}
