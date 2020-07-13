package models

import "database/sql"

type ConfigMap struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}
