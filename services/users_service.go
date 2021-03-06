package services

import (
	"github.com/angadthandi/bookstore_users-api/domain/users"
	"github.com/angadthandi/bookstore_users-api/utils/crypto_utils"
	"github.com/angadthandi/bookstore_users-api/utils/date_utils"
	"github.com/angadthandi/bookstore_utils-go/rest_errors"
)

type usersServiceInterface interface {
	GetUser(int64) (*users.User, rest_errors.RestErr)
	CreateUser(users.User) (*users.User, rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, rest_errors.RestErr)
	DeleteUser(int64) rest_errors.RestErr
	Search(string) (users.Users, rest_errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, rest_errors.RestErr)
}

type usersService struct{}

func NewUserService() usersServiceInterface {
	return &usersService{}
}

func (usrvc *usersService) GetUser(id int64) (*users.User, rest_errors.RestErr) {
	ret := &users.User{ID: id}
	err := ret.Get()
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (usrvc *usersService) CreateUser(u users.User) (*users.User, rest_errors.RestErr) {
	err := u.Validate()
	if err != nil {
		return nil, err
	}

	u.DateCreated = date_utils.GetNowDBFormat()
	u.Status = users.StatusActive
	u.Password = crypto_utils.GetMD5(u.Password)
	err = u.Save()
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (usrvc *usersService) UpdateUser(isPartial bool, u users.User) (*users.User, rest_errors.RestErr) {
	currUser, err := usrvc.GetUser(u.ID)
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

func (usrvc *usersService) DeleteUser(id int64) rest_errors.RestErr {
	user := &users.User{ID: id}
	return user.Delete()
}

func (usrvc *usersService) Search(status string) (users.Users, rest_errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (usrvc *usersService) LoginUser(
	r users.LoginRequest,
) (*users.User, rest_errors.RestErr) {
	ret := &users.User{
		Email:    r.Email,
		Password: crypto_utils.GetMD5(r.Password),
	}
	err := ret.FindByEmailAndPassword()
	if err != nil {
		return nil, err
	}

	return ret, nil
}
