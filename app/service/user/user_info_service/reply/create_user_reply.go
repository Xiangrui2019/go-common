package reply

import (
	"go-common/app/service/user/user_info_service/dtos"
	"go-common/library/net/rpc/reply"
)

type CreateUserReply struct {
	reply.Reply
	Data *dtos.UserDto
}
