package ecode

import (
	"github.com/pkg/errors"
)

var (
	// 数据库错误
	DatabaseRecordNotFound = errors.New("Database Record Not Found")
	DatabaseRecordExist    = errors.New("Database Record is Exist")
)
