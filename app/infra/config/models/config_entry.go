package models

type ConfigEntry struct {
	Id          int64  `json:"id"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	ConfigMapID int64  `json:"config_map_id"`
}
