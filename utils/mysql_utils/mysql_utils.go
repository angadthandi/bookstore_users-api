package mysql_utils

import (
	"errors"
	"fmt"

	"github.com/angadthandi/bookstore_utils-go/rest_errors"
	"github.com/go-sql-driver/mysql"
)

func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return rest_errors.NewInternalServerError(
			"error processing request",
			errors.New(fmt.Sprintf("database error: %v", err.Error())),
		)
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("invalid data")
	}

	return rest_errors.NewInternalServerError(
		"error processing request",
		errors.New(fmt.Sprintf("database error: %v", err.Error())),
	)
}
