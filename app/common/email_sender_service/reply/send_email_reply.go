package reply

import "go-common/library/ecode"

type SendEmailReply struct {
	ErrorCode ecode.Code
	Message   string
	Details   string
}
