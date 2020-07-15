package reply

import (
	"go-common/library/net/rpc/reply"
	"time"
)

type GetConfigEntrysByMapIdReply struct {
	reply.Reply
	Datas []*GetConfigEntryDto
}

type GetConfigEntryDto struct {
	ID          uint
	Key         string
	Value       string
	ConfigMapId uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
