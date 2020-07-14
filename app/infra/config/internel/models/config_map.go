package models

import "github.com/jinzhu/gorm"

type ConfigMap struct {
	gorm.Model
	Name        string
	Description string
}
