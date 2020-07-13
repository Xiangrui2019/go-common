package ecode

import (
	"strconv"
)

type Code int64

func New(e int64) Code {
	return Code(e)
}

func (e Code) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

func (e Code) Equal(code Code) bool {
	return code == e
}
