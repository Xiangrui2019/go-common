package services

import (
	"context"
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
	reqMetadata := ctx.Value(share.ReqMetaDataKey).(map[string]string)

	if reqMetadata["token"] != conf.String("CONFIGURE_TOKEN", "") {
		rep.ErrorCode = ecode.UnauthorizedException
		rep.Message = "The client token is not vaild."
		rep.Details = "客户端token错误."

		return errors.New("The client token is not vaild.")
	}

	return nil
}
