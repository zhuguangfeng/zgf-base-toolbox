package utils

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func IsDuplicateKeyError(err error) bool {
	var errMysql *mysql.MySQLError
	if errors.As(err, &errMysql) && errMysql.Number == 1062 {
		return true
	}
	return false
}
func IsRecordNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
