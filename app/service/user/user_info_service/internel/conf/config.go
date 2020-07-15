package conf

import (
	"go-common/app/service/user/user_info_service/internel/database/orm"
	"go-common/library/conf"

	"github.com/joho/godotenv"
)

func Init() error {
	godotenv.Load()

	// 链接MySQL数据库
	err := orm.ConnectDatabase(conf.String("DATABASE_DSN", ""), conf.Int("DATABASE_TIMEOUT", 15))

	if err != nil {
		return err
	}

	return nil
}
