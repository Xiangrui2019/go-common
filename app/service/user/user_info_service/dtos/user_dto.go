package dtos

import (
	"go-common/app/service/user/user_info_service/internel/models"
	"time"
)

type UserDto struct {
	ID                   uint
	NikeName             string
	UserName             string
	Bio                  string
	Sex                  int64
	Avatar               string
	Email                string
	EmailConfirmed       bool
	PhoneNumber          string
	PhoneNumberConfirmed bool
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

// Map *models.User to *dto.UserDto Model
func AutoMapperUserDto(user *models.User) *UserDto {
	return &UserDto{
		ID:                   user.ID,
		NikeName:             user.NikeName,
		UserName:             user.UserName,
		Bio:                  user.Bio,
		Sex:                  user.Sex,
		Avatar:               user.Avatar,
		Email:                user.Email,
		EmailConfirmed:       user.EmailConfirmed,
		PhoneNumber:          user.PhoneNumber,
		PhoneNumberConfirmed: user.PhoneNumberConfirmed,
		CreatedAt:            user.CreatedAt,
		UpdatedAt:            user.UpdatedAt,
	}
}
