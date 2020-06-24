package services

import (
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"github.com/nao4869/golang-bookstore-user-api/domain/users"
)

// GetUser -
func GetUser(userID int64) (*users.User, rest_errors.RestErr) {
	result := &users.User{ID: userID}
	if error := result.Get(); error != nil {
		return nil, error
	}
	return result, nil
}

// CreateUser - must not return both users and errors - only return one
// if we return nil, nil - this will cause many error in users_controllers when handling nil error and returing nil result etc
func CreateUser(user users.User) (*users.User, rest_errors.RestErr) {
	if error := user.Validate(); error != nil {
		return nil, error
	}

	if error := user.Save(); error != nil {
		return nil, error
	}
	return &user, nil
}
