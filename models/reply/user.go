/*
 * @Author: lihuan
 * @Date: 2021-10-14 11:32:22
 * @LastEditTime: 2021-10-14 13:06:29
 * @Email: 17719495105@163.com
 */
package reply

type UserReply struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Sex      int32  `json:"sex"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Hobby    string `json:"hobby"`
}
