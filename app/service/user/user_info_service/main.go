package main

import (
	"go-common/app/service/user/user_info_service/internel/conf"
	"log"
)

func main() {
	err := conf.Init()
	if err != nil {
		log.Fatal(err)
	}
}
