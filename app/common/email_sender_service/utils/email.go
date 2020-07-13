package utils

import (
	"go-common/app/common/email_sender_service/request"
	"net/smtp"
	"os"
	"strings"
)

// 最简单的发送邮件的基本方法
func SendToMail(to, subject, body string, mail_type int64) error {
	var content_type string

	hp := strings.Split(os.Getenv("STMP_HOST"), ":")
	auth := smtp.PlainAuth("", os.Getenv("STMP_USERNAME"), os.Getenv("STMP_PASSWORD"), hp[0])

	if mail_type == int64(request.HTML) {
		content_type = "Content-Type: text/html" + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + os.Getenv("STMP_USERNAME") + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(os.Getenv("STMP_HOST"), auth, os.Getenv("STMP_USERNAME"), send_to, msg)
	return err
}
