package dao

import (
	"go-common/app/infra/config/database"
	"go-common/app/infra/config/models"

	"github.com/pkg/errors"
)

func CreateConfigEntry(config_entry *models.ConfigEntry) (int64, error) {
	sql := `INSERT t_config_entry (key, value, config_map_id) VALUES (:key, :value, :config_map_id)`
	result, err := database.Database.NamedExec(sql, config_entry)
	if err != nil {
		return 0, errors.Wrap(err, "exec sql language error.")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "insert data error.")
	}

	return id, nil
}

func ConfigEntrys(config_map_id int64) ([]models.ConfigEntry, error) {
	var config_entrys []models.ConfigEntry
	sql := `SELECT * FROM t_config_entry WHERE id=?`

	err := database.Database.Select(&config_entrys, sql, config_entrys)
	if err != nil {
		return nil, errors.Wrap(err, "get configentrys error.")
	}

	return config_entrys, nil
}

func UpdateConfigEntry(config_entry *models.ConfigEntry) error {
	sql := `UPDATE t_config_entry SET key=:key value=:value config_map_id=:config_map_id WHERE id=:id`
	_, err := database.Database.NamedExec(sql, config_entry)

	if err != nil {
		return errors.Wrap(err, "update configentry error.")
	}

	return nil
}

func DeleteConfigEntry(config_map_id int64) error {
	sql := `DELETE FROM t_config_entry WHERE id=?`
	_, err := database.Database.Exec(sql, config_map_id)

	if err != nil {
		return errors.Wrap(err, "delete config_entrys error.")
	}

	return nil
}
