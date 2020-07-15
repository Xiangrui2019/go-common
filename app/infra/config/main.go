package main

import (
	"go-common/app/infra/config/internel/conf"
	"go-common/app/infra/config/internel/services"
	cfg "go-common/library/conf"
	"log"

	"github.com/smallnest/rpcx/server"
)

func main() {
	err := conf.Init()

	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer()
	s.Register(new(services.ConfigMapService), "")
	s.Register(new(services.ConfigEntryService), "")
	s.Register(new(services.ConfigingService), "")
	s.Serve("tcp", cfg.String("RPC_ADDR", ":8848"))
}
