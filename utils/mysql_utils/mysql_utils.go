package mysql_utils

import (
	"fmt"

	"github.com/angadthandi/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return errors.NewInternalServerError(
			fmt.Sprintf("database error: %v", err.Error()),
		)
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}

	return errors.NewInternalServerError(
		fmt.Sprintf("database error: %v", err.Error()),
	)
}
