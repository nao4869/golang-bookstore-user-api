package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/nao4869/golang-bookstore-user-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

// ParseError -
func ParseError(err error) *errors.RestError {
	sqlError, okay := err.(*mysql.MySQLError)
	// in case unable to convert to MySQLError
	if !okay {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlError.Number {
	case 1062:
		return errors.NewBadRequestError("duplicated key")
	}
	return errors.NewInternalServerError("error processing request")
}
