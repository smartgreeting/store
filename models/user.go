/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-13 09:47:02
 * @Email: 17719495105@163.com
 */
package models

type User struct {
	ID        uint64
	Username  string
	Password  string
	Avatar    string
	Sex       int
	Phone     string
	Email     string
	Address   string
	Hobby     string
	Deleted   int
	CreatedAt int
	UpdatedAt int
}

func (m User) TableName() string {
	return "hc_user"
}
