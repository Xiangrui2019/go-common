package dao

import (
	"go-common/app/infra/config/internel/database"
	"go-common/app/infra/config/internel/models"

	"github.com/pkg/errors"
)

// 创建ConfigMap
func CreateConfigMap(config_map *models.ConfigMap) (*models.ConfigMap, error) {
	err := database.DB.Create(config_map).Error

	if err != nil {
		return nil, errors.Wrap(err, "Create config_maps had an exception.")
	}

	return config_map, nil
}

// 获取ConfigMaps
func ConfigMaps() ([]*models.ConfigMap, error) {
	configmaps := []*models.ConfigMap{}

	if err := database.DB.Find(&configmaps).Error; err != nil {
		return nil, errors.Wrap(err, "Getting all config_maps had an exception.")
	}

	return configmaps, nil
}

// 通过id获取一个ConfigMap
func ConfigMap(id uint) (*models.ConfigMap, error) {
	configmap := models.ConfigMap{}

	if err := database.DB.First(&configmap).Error; err != nil {
		return nil, errors.Wrap(err, "Getting a config_map had an exception.")
	}

	return &configmap, nil
}

// 更新ConfigMap
func UpdateConfigMap(config_map *models.ConfigMap) error {
	if err := database.DB.Save(&config_map).Error; err != nil {
		return errors.Wrap(err, "Update a config_map had an exception.")
	}

	return nil
}

// 删除ConfigMap
func DeleteConfigMap(id uint) error {
	config_map := models.ConfigMap{}

	err := database.DB.First(&config_map, id).Error
	if err != nil {
		return errors.Wrap(err, "Delete a config_map had an exception, the configmap is not found.")
	}

	err = database.DB.Delete(&config_map).Error
	if err != nil {
		return errors.Wrap(err, "Delete a config_map had an exception.")
	}

	return nil
}
