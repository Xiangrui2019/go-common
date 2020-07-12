package main

import (
	_ "go-common/app/infra/config/conf"
	"go-common/app/infra/config/servers"
	"log"
	"os"
)

func main() {
	server, err := servers.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	err = server.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(os.Getenv("HTTP_ADDR"))
	if err != nil {
		log.Fatal(err)
	}
}
