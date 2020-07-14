package reply

import "go-common/library/net/rpc/reply"

type ConfigureReply struct {
	reply.Reply
	Entrys []*ConfigureEntryDto
}

type ConfigureEntryDto struct {
	Key   string
	Value string
}
