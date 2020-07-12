package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var Database *sqlx.DB

func ConnectDatabase(database_dsn string) error {
	var err error
	Database, err = sqlx.Open("mysql", database_dsn)

	if err != nil {
		return errors.Wrap(err, "Database connection creating error. Please check you network.")
	}

	if err := Database.Ping(); err != nil {
		return errors.Wrap(err, "Database connection ping error. Please check you database connection string.")
	}

	return nil
}
