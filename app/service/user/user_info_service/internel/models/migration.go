package models

func MigrationDatabase() {
	// 迁移数据库, User表
	GormEngine.AutoMigrate(&User{})
}
