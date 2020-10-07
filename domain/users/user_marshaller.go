package users

import "encoding/json"

type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func (users Users) Marshal(isPublic bool) ([]interface{}, error) {
	var err error

	ret := make([]interface{}, len(users))
	for i, user := range users {
		ret[i], err = user.Marshal(isPublic)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (user *User) Marshal(isPublic bool) (interface{}, error) {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	if isPublic {
		var public PublicUser
		err = json.Unmarshal(userJSON, &public)
		if err != nil {
			return nil, err
		}

		return public, nil
	}

	var private PrivateUser
	err = json.Unmarshal(userJSON, &private)
	if err != nil {
		return nil, err
	}

	return private, nil
}
