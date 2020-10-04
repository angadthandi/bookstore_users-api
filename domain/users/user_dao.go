package users

import (
	"fmt"
	"log"

	"github.com/angadthandi/bookstore_users-api/datasources/mysql/users_db"
	"github.com/angadthandi/bookstore_users-api/utils/date_utils"
	"github.com/angadthandi/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Get() *errors.RestErr {
	err := users_db.Client.Ping()
	if err != nil {
		log.Fatalf("unable to connect to mysql db error: %v", err)
	}

	ret := usersDB[u.ID]
	if ret == nil {
		return errors.NewNotFoundError(
			fmt.Sprintf("user %d not found", u.ID),
		)
	}

	u.ID = ret.ID
	u.FirstName = ret.FirstName
	u.LastName = ret.LastName
	u.Email = ret.Email
	u.DateCreated = ret.DateCreated

	return nil
}

func (u *User) Save() *errors.RestErr {
	curr := usersDB[u.ID]
	if curr != nil {
		if curr.Email == u.Email {
			return errors.NewBadRequestError(
				fmt.Sprintf("email %s already registered", u.Email),
			)
		}
		return errors.NewBadRequestError(
			fmt.Sprintf("user %d already exists", u.ID),
		)
	}

	u.DateCreated = date_utils.GetNowString()

	usersDB[u.ID] = u

	return nil
}
