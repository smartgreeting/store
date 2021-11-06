/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-13 09:47:02
 * @Email: 17719495105@163.com
 */
package models

import "store/rpc-user/apiuser"

type User struct {
	ID        uint64
	Username  string
	Password  string `json:"-"`
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

func RpcToUser(in *apiuser.UserReply) User {
	return User{
		ID : in.Id,
		Username : in.Username,
		Avatar : in.Avatar,
		Sex : int(in.Sex),
		Phone : in.Phone,
		Email : in.Email,
		Address : in.Address,
		Hobby : in.Hobby,
	}
	
	
}

func (u User) TableName() string {
	return "hc_user"
}
