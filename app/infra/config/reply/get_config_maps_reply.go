package reply

import (
	"go-common/library/net/rpc/reply"
	"time"
)

type GetConfigMapsReply struct {
	reply.Reply
	Data []*GetConfigMapDto
}

type GetConfigMapDto struct {
	ID          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
