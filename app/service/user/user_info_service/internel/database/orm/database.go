package orm

import (
	"go-common/library/ecode"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var GormEngine *gorm.DB

func init() {
	gorm.ErrRecordNotFound = ecode.DatabaseRecordNotFound
}

func ConnectDatabase(connectionString string, timeout int) error {
	var err error

	GormEngine, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return errors.Wrap(err, "Gorm engine creating had an exception.")
	}

	MigrationDatabase()

	return nil
}
