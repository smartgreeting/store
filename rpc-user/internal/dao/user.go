/*
 * @Author: lihuan
 * @Date: 2021-10-11 08:43:38
 * @LastEditTime: 2021-10-14 17:28:52
 * @Email: 17719495105@163.com
 */
package dao

import (
	"context"
	"store/models"

	"gorm.io/gorm"
)

type UserDao struct {
	ctx context.Context
	db  *gorm.DB
}

func NewUserDao(ctx context.Context, db *gorm.DB) *UserDao {
	return &UserDao{
		ctx: ctx,
		db:  db,
	}
}

func (d UserDao) Create(user *models.User) (*models.User, error) {
	err := d.db.Create(user).Error
	return user, err
}

func (d UserDao) FindByPhone(phone string) (*models.User, error) {
	var user models.User

	err := d.db.Where("phone = ?", phone).First(&user).Error

	return &user, err
}
