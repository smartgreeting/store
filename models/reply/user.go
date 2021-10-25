/*
 * @Author: lihuan
 * @Date: 2021-10-14 11:32:22
 * @LastEditTime: 2021-10-25 15:00:04
 * @Email: 17719495105@163.com
 */
package reply

type UserReply struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Sex      int32  `json:"sex"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Hobby    string `json:"hobby"`
}
