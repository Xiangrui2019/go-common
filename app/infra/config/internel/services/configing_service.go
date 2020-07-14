package services

import (
	"context"
	"go-common/app/infra/config/internel/dao"
	"go-common/app/infra/config/reply"
	"go-common/app/infra/config/request"
	"go-common/library/conf"
	"go-common/library/ecode"

	"github.com/pkg/errors"

	"github.com/smallnest/rpcx/share"
)

type ConfigingService struct {
}

func (service *ConfigingService) Configure(ctx context.Context, args *request.ConfigureRequest, rep *reply.ConfigureReply) error {
	var reply_dto []*reply.ConfigureEntryDto

	reqMetadata := ctx.Value(share.ReqMetaDataKey).(map[string]string)

	if reqMetadata["token"] != conf.String("CONFIGURE_TOKEN", "") {
		rep.ErrorCode = ecode.UnauthorizedException
		rep.Message = "The client token is not vaild."
		rep.Details = "客户端token错误."
		rep.Entrys = nil

		return errors.New("The client token is not vaild.")
	}

	entrys, err := dao.ConfigEntrysByMapId(args.ConfigMapId)
	if err != nil {
		rep.ErrorCode = ecode.ServerException
		rep.Message = "The config entrys getting had an exception."
		rep.Details = "配置表获取时有一个错误, 无法获取."
		rep.Entrys = nil
		return err
	}

	for _, v := range entrys {
		reply_dto = append(reply_dto, &reply.ConfigureEntryDto{
			Key:   v.Key,
			Value: v.Value,
		})
	}

	rep.ErrorCode = ecode.OK
	rep.Message = "The config entrys getting ok."
	rep.Details = ""
	rep.Entrys = reply_dto

	return nil
}
