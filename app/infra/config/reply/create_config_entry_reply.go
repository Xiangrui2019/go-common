package reply

import (
	"go-common/library/net/rpc/reply"
	"time"
)

type CreateConfigEntryReply struct {
	reply.Reply
	Data *CreateConfigEntryDto
}

type CreateConfigEntryDto struct {
	ID          uint
	Key         string
	Value       string
	ConfigMapId uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
