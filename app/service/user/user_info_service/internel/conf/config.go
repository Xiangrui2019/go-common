package conf

import (
	"go-common/app/service/user/user_info_service/internel/models"
	"go-common/library/conf"

	"github.com/joho/godotenv"
)

func Init() error {
	godotenv.Load()

	// 链接MySQL数据库
	err := models.ConnectDatabase(conf.String("DATABASE_DSN", ""), conf.Int("DATABASE_TIMEOUT", 15))

	if err != nil {
		return err
	}

	return nil
}
