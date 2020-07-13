package reply

type DefaultReply struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func (reply *DefaultReply) GetMessage() string {
	return reply.Message
}

func (reply *DefaultReply) GetDetails() string {
	return reply.Details
}
