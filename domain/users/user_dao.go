package users

import (
	"errors"
	"fmt"

	"github.com/angadthandi/bookstore_users-api/datasources/mysql/users_db"
	// "github.com/angadthandi/bookstore_users-api/logger"
	"github.com/angadthandi/bookstore_utils-go/logger"
	"github.com/angadthandi/bookstore_utils-go/rest_errors"
)

const (
	queryInsertUser = `INSERT INTO users
		(first_name, last_name, email, date_created, status, password)
		VALUES (?, ?, ?, ?, ?, ?)`
	queryGetUser = `SELECT id, first_name, last_name, email, date_created, status
		FROM users WHERE id=?`
	queryUpdateUser = `UPDATE users SET first_name=?, last_name=?, email=?
		WHERE id=?`
	queryDeleteUser   = `DELETE FROM users WHERE id=?`
	queryFindByStatus = `SELECT id, first_name, last_name, email, date_created, status
		FROM users WHERE status=?`
	queryFindByEmailAndPassword = `SELECT id, first_name, last_name, email, date_created, status
		FROM users WHERE email=? AND password=? AND status=?`
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Get() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError(
			"error when trying to get user",
			errors.New("database error"),
		)
	}
	defer stmt.Close()

	ret := stmt.QueryRow(u.ID)
	err = ret.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.DateCreated,
		&u.Status,
	)
	if err != nil {
		logger.Error("error when trying to get user by id", err)
		return rest_errors.NewInternalServerError(
			"error when trying to get user",
			errors.New("database error"),
		)
		// return mysql_utils.ParseError(err)
	}

	// ret := usersDB[u.ID]
	// if ret == nil {
	// 	return rest_errors.NewNotFoundError(
	// 		fmt.Sprintf("user %d not found", u.ID),
	// 	)
	// }

	// u.ID = ret.ID
	// u.FirstName = ret.FirstName
	// u.LastName = ret.LastName
	// u.Email = ret.Email
	// u.DateCreated = ret.DateCreated

	return nil
}

func (u *User) Save() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return rest_errors.NewInternalServerError(
			"error when trying to save user",
			errors.New("database error"),
		)
		// return rest_errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertRet, err := stmt.Exec(
		u.FirstName, u.LastName, u.Email, u.DateCreated, u.Status, u.Password,
	)
	if err != nil {
		logger.Error("error when trying to save user", err)
		return rest_errors.NewInternalServerError(
			"error when trying to save user",
			errors.New("database error"),
		)
		// return mysql_utils.ParseError(err)
	}

	userID, err := insertRet.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert user id", err)
		return rest_errors.NewInternalServerError(
			"error when trying to save user",
			errors.New("database error"),
		)
		// return mysql_utils.ParseError(err)
	}

	u.ID = userID

	// curr := usersDB[u.ID]
	// if curr != nil {
	// 	if curr.Email == u.Email {
	// 		return rest_errors.NewBadRequestError(
	// 			fmt.Sprintf("email %s already registered", u.Email),
	// 		)
	// 	}
	// 	return rest_errors.NewBadRequestError(
	// 		fmt.Sprintf("user %d already exists", u.ID),
	// 	)
	// }

	// u.DateCreated = date_utils.GetNowString()

	// usersDB[u.ID] = u

	return nil
}

func (u *User) Update() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return rest_errors.NewInternalServerError(
			"error when trying to update user",
			errors.New("database error"),
		)
		// return rest_errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		u.FirstName, u.LastName, u.Email, u.ID,
	)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return rest_errors.NewInternalServerError(
			"error when trying to update user",
			errors.New("database error"),
		)
		// return mysql_utils.ParseError(err)
	}

	return nil
}

func (u *User) Delete() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return rest_errors.NewInternalServerError(
			"error when trying to delete user",
			errors.New("database error"),
		)
		// return rest_errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID)
	if err != nil {
		logger.Error("error when trying to delete user", err)
		return rest_errors.NewInternalServerError(
			"error when trying to delete user",
			errors.New("database error"),
		)
		// return mysql_utils.ParseError(err)
	}

	return nil
}

func (u *User) FindByStatus(status string) ([]User, rest_errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare findbystatus user statement", err)
		return nil, rest_errors.NewInternalServerError(
			"error when trying to get user",
			errors.New("database error"),
		)
		// return nil, rest_errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to findbystatus user", err)
		return nil, rest_errors.NewInternalServerError(
			"error when trying to get user",
			errors.New("database error"),
		)
		// return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()

	var ret []User
	for rows.Next() {
		var user User
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.DateCreated,
			&user.Status,
		)
		if err != nil {
			logger.Error("error when trying to scan rows for findbystatus user", err)
			return nil, rest_errors.NewInternalServerError(
				"error when trying to get user",
				errors.New("database error"),
			)
			// return nil, mysql_utils.ParseError(err)
		}

		ret = append(ret, user)
	}

	if len(ret) == 0 {
		return nil, rest_errors.NewNotFoundError(
			fmt.Sprintf("no user matching status: %v", status),
		)
	}

	return ret, nil
}

func (u *User) FindByEmailAndPassword() rest_errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError(
			"error when trying to get user",
			errors.New("database error"),
		)
	}
	defer stmt.Close()

	ret := stmt.QueryRow(u.Email, u.Password, StatusActive)
	err = ret.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.DateCreated,
		&u.Status,
	)
	if err != nil {
		logger.Error("error when trying to get user by email and password", err)
		return rest_errors.NewInternalServerError(
			"error when trying to get user",
			errors.New("database error"),
		)
	}

	return nil
}
