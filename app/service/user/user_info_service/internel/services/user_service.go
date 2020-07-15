package services

import (
	"context"
	"go-common/app/service/user/user_info_service/dtos"
	"go-common/app/service/user/user_info_service/internel/models"
	"go-common/app/service/user/user_info_service/reply"
	"go-common/app/service/user/user_info_service/request"
	"go-common/library/conf"
	"go-common/library/ecode"
)

type UserService struct {
}

func (service *UserService) CreateUser(ctx context.Context, req *request.CreateUserRequest, repl *reply.CreateUserReply) error {
	user := &models.User{
		NikeName:             req.NikeName,
		UserName:             req.UserName,
		Bio:                  "这个人很懒, 什么都没写.....",
		Sex:                  int64(0),
		Avatar:               conf.String("DEFAULT_AVATAR", ""),
		Email:                req.Email,
		EmailConfirmed:       false,
		PhoneNumber:          req.PhoneNumber,
		PhoneNumberConfirmed: false,
	}

	u, err := models.CreateUser(user)
	if err != nil {
		repl.ErrorCode = ecode.ServerException
		repl.Message = "用户创建失败"
		repl.Details = "服务器开小差了."
		repl.Data = nil

		return err
	}

	repl.ErrorCode = ecode.OK
	repl.Message = "OK"
	repl.Details = "用户创建成功."
	repl.Data = dtos.AutoMapperUserDto(u)

	return nil
}
