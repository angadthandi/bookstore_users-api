package users

import (
	"strings"

	"github.com/angadthandi/bookstore_utils-go/rest_errors"
)

const (
	StatusActive = "active"
)

type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (u *User) Validate() rest_errors.RestErr {
	u.FirstName = strings.TrimSpace(strings.ToLower(u.FirstName))
	u.LastName = strings.TrimSpace(strings.ToLower(u.LastName))
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return rest_errors.NewBadRequestError("invalid email")
	}

	u.Password = strings.TrimSpace(u.Password)
	if u.Password == "" {
		return rest_errors.NewBadRequestError("invalid password")
	}
	return nil
}
