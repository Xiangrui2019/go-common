package dao

import (
	"go-common/app/infra/config/internel/database"
	"go-common/app/infra/config/internel/models"

	"github.com/pkg/errors"
)

// 创建配置项
func CreateConfigEntry(config_entry *models.ConfigEntry) (*models.ConfigEntry, error) {
	err := database.DB.Create(config_entry).Error
	if err != nil {
		return nil, errors.Wrap(err, "Create config_entry had an exception.")
	}

	return config_entry, nil
}

// 查看配置项(by map_id)
func ConfigEntrysByMapId(configmap_id uint) ([]*models.ConfigEntry, error) {
	configentrys := []*models.ConfigEntry{}

	if err := database.DB.Find(&configentrys).Error; err != nil {
		return nil, errors.Wrap(err, "Showing Config Entrys by configmapid had an exception.")
	}

	return configentrys, nil
}

// 更新配置项
func UpdateConfigEntry(config_entry *models.ConfigEntry) error {
	if err := database.DB.Save(&config_entry).Error; err != nil {
		return errors.Wrap(err, "Update a config_entry had an exception.")
	}

	return nil
}

// 删除配置项
func DeleteConfigEntry(id uint) error {
	config_entry := models.ConfigEntry{}

	err := database.DB.First(&config_entry, id).Error
	if err != nil {
		return errors.Wrap(err, "Delete a config_entry had an exception, the config_entry is not found.")
	}

	err = database.DB.Delete(&config_entry).Error
	if err != nil {
		return errors.Wrap(err, "Delete a config_entry had an exception.")
	}

	return nil
}
