package models

import "github.com/jinzhu/gorm"

type ConfigEntry struct {
	gorm.Model
	Key         string
	Value       string
	ConfigMapId uint
}
