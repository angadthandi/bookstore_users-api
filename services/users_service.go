package services

import (
	"github.com/angadthandi/bookstore_users-api/domain/users"
	"github.com/angadthandi/bookstore_users-api/utils/errors"
)

func GetUser(id int64) (*users.User, *errors.RestErr) {
	ret := &users.User{ID: id}
	err := ret.Get()
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func CreateUser(u users.User) (*users.User, *errors.RestErr) {
	err := u.Validate()
	if err != nil {
		return nil, err
	}

	err = u.Save()
	if err != nil {
		return nil, err
	}

	return &u, nil
}
