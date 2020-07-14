package database

import (
	"go-common/app/infra/config/internel/models"
)

func MigrationDatabase() {
	DB.AutoMigrate(&models.ConfigMap{})
}
