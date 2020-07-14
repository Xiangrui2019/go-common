package conf

import (
	"go-common/app/infra/config/internel/database"
	"go-common/library/conf"

	"github.com/joho/godotenv"
)

func Init() error {
	godotenv.Load()

	err := database.ConnectDatabase(conf.String("DATABASE_DSN", ""), int64(conf.Int("DATABASE_TIMEOUT", 10)))
	if err != nil {
		return err
	}

	return nil
}
