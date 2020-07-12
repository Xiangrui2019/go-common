package conf

import (
	"go-common/app/infra/config/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	err := database.ConnectDatabase(os.Getenv("DATABASE_DSN"))
	if err != nil {
		log.Fatal(err)
	}
}
