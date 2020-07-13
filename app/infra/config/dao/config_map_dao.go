package dao

import (
	"go-common/app/infra/config/database"
	"go-common/app/infra/config/models"

	"github.com/pkg/errors"
)

// 创建配置表
func CreateConfigMap(config_map *models.ConfigMap) (int64, error) {
	sql := `INSERT INTO t_config_map (name, description) VALUES(:name,:description)`
	result, err := database.Database.NamedExec(sql, config_map)

	if err != nil {
		return 0, errors.Wrap(err, "exec sql language error.")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "insert data error.")
	}

	return id, nil
}

// 删除配置表
func DeleteConfigMap(id int64) error {
	sql := `DELETE FROM t_config_map WHERE id=?`
	_, err := database.Database.Exec(sql, id)

	if err != nil {
		return errors.Wrap(err, "delete configmap error.")
	}

	return nil
}

// 更新配置表
func UpdateConfigMap(config_map *models.ConfigMap) error {
	sql := `UPDATE t_config_map SET name=:name description=:description WHERE id=:id`
	_, err := database.Database.NamedExec(sql, config_map)

	if err != nil {
		return errors.Wrap(err, "update configmap error.")
	}

	return nil
}

// 查看配置表
func ConfigMap(id int64) (*models.ConfigMap, error) {
	var config_map models.ConfigMap
	sql := `SELECT * FROM t_config_map WHERE id=?`
	err := database.Database.Get(&config_map, sql, id)
	if err != nil {
		return nil, errors.Wrap(err, "get config_map errors")
	}

	return &config_map, nil
}

// 列出配置表
func ConfigMaps() (*[]models.ConfigMap, error) {
	var config_maps []models.ConfigMap
	sql := `SELECT * FROM t_config_map`

	err := database.Database.Select(&config_maps, sql)
	if err != nil {
		return nil, errors.Wrap(err, "get configmaps error.")
	}

	return &config_maps, nil
}
