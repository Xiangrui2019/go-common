package reply

import (
	"go-common/library/net/rpc/reply"
	"time"
)

type CreateConfigMapReply struct {
	reply.Reply
	Data *CreateConfigMapDto
}

type CreateConfigMapDto struct {
	ID          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
