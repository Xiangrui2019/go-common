package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// 用户数据库模型
type User struct {
	gorm.Model
	// 昵称
	NikeName string
	// 用户名
	Username string
	// 简介
	Bio string
	// 性别([x]迪诺)
	Sex int64
	// 头像
	Avatar string
	// 电子邮件
	Email string
	// 电子邮件确认
	EmailConfirmed bool
	// 手机号
	PhoneNumber string
	// 手机号确认
	PhoneNumberConfirmed bool
}

func CreateUser(user *User) (*User, error) {
	err := GormEngine.Create(user).Error

	if err != nil {
		return nil, errors.Wrap(err, "Creating user had an unknown exception.")
	}

	return user, nil
}

func GetUser(id uint) (*User, error) {
	var user User

	err := GormEngine.First(&user, id).Error
	if err != nil {
		return nil, errors.Wrap(err, "Getting user had an exception.")
	}

	return &user, nil
}

func GetUsers() ([]*User, error) {
	var users []*User

	err := GormEngine.Find(&users).Error
	if err != nil {
		return nil, errors.Wrap(err, "Getting users had an exception.")
	}

	return users, nil
}

func UpdateUser(user *User) error {
	err := GormEngine.Save(user).Error
	if err != nil {
		return errors.Wrap(err, "Updating user had an exception.")
	}

	return nil
}

func DeleteUser(user *User) error {
	err := GormEngine.Delete(user).Error
	if err != nil {
		return errors.Wrap(err, "Deleting user had an exception.")
	}

	return nil
}
