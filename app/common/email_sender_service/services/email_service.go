package services

import (
	"context"
	"go-common/app/common/email_sender_service/reply"
	"go-common/app/common/email_sender_service/request"
	"go-common/app/common/email_sender_service/utils"
	"go-common/library/ecode"

	"github.com/pkg/errors"
)

type EmailService struct{}

func (service *EmailService) SendEmail(context context.Context, args *request.SendEmailRequest, reply *reply.SendEmailReply) error {
	err := utils.SendToMail(args.To, args.Subject, args.Body, int64(args.Type))

	if err != nil {
		reply.ErrorCode = ecode.ServerException
		reply.Message = "发送电子邮件失败, 请检查配置文件"
		reply.Details = "服务器开小差了."

		return errors.Wrap(err, "发送电子邮件失败, 请检查配置文件.")
	}

	reply.ErrorCode = ecode.OK
	reply.Message = "成功发送电子邮件."
	reply.Details = "成功发送电子邮件."

	return nil
}
