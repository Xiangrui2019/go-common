package services

import (
	"context"
	"go-common/app/common/email_sender_service/reply"
	"go-common/app/common/email_sender_service/request"
	"go-common/app/common/email_sender_service/utils"
)

type EmailService struct {
}

func (service *EmailService) SendEmail(context context.Context, args *request.SendEmailRequest, reply *reply.SendEmailReply) error {
	utils.SendToMail(args.To, args.Subject, args.Body, args.Type)
}
