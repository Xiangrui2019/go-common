package reply

import "go-common/library/ecode"

// 标准RPC返回
type Reply struct {
	ErrorCode ecode.Code `json:"error_code"`
	Message   string     `json:"message,omitempty"`
	Details   string     `json:"details,omitempty"`
}
