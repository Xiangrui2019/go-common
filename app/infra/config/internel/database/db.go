package database

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
)

var DB *gorm.DB

func ConnectDatabase(connectionString string, timeout int64) error {
	var err error

	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return errors.Wrap(err, "创建数据库连接失败, 请检查您的网络和程序.")
	}

	DB.DB().SetMaxIdleConns(50)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(time.Second * time.Duration(timeout))

	MigrationDatabase()

	return nil
}
