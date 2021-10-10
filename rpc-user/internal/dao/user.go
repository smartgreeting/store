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
	err := d.db.Where("phone = ?", phone).Find(&user).Error
	return &user, err
}
