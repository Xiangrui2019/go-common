package main

import (
	"go-common/app/infra/config/internel/conf"
	"log"
)

func main() {
	err := conf.Init()

	if err != nil {
		log.Fatal(err)
	}
}
