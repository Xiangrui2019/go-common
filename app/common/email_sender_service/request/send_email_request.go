package request

type EmailType int64

var (
	HTML  EmailType = EmailType(0)
	PLAIN EmailType = EmailType(1)
)

type SendEmailRequest struct {
	To      string
	Subject string
	Body    string
	Type    EmailType
}
