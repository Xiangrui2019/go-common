package conf

import (
	"go-common/app/infra/config/database"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	database.ConnectDatabase()
}
