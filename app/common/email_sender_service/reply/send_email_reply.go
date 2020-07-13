package reply

import "go-common/library/ecode"

type SendEmailReply struct {
	ErrorCode ecode.Code `json:"err_code"`
	Message   string     `json:"message"`
	Details   string     `json:"details"`
}
