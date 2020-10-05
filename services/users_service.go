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

func UpdateUser(isPartial bool, u users.User) (*users.User, *errors.RestErr) {
	currUser, err := GetUser(u.ID)
	if err != nil {
		return nil, err
	}

	// err = u.Validate()
	// if err != nil {
	// 	return nil, err
	// }

	if isPartial {
		if u.FirstName != "" {
			currUser.FirstName = u.FirstName
		}
		if u.LastName != "" {
			currUser.LastName = u.LastName
		}
		if u.Email != "" {
			currUser.Email = u.Email
		}
	} else {
		currUser.FirstName = u.FirstName
		currUser.LastName = u.LastName
		currUser.Email = u.Email
	}

	err = currUser.Update()
	if err != nil {
		return nil, err
	}

	return currUser, nil
}

func DeleteUser(id int64) *errors.RestErr {
	user := &users.User{ID: id}
	return user.Delete()
}
