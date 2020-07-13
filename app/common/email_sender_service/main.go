package main

import (
	"go-common/app/common/email_sender_service/conf"
	"go-common/app/common/email_sender_service/services"

	shared_conf "go-common/library/conf"

	"github.com/smallnest/rpcx/server"
)

func main() {
	conf.Init()

	s := server.NewServer()
	s.Register(new(services.EmailService), "")
	s.Serve("tcp", shared_conf.String("RPC_ADDR", ":7015"))
}
