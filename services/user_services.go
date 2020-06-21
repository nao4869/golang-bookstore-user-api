package services

import (
	"github.com/nao4869/golang-bookstore-user-api/domain/users"
	"github.com/nao4869/golang-bookstore-user-api/utils/errors"
)

// CreateUser - must not return both users and errors - only return one
// if we return nil, nil - this will cause many error in users_controllers when handling nil error and returing nil result etc
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if error := users.Validate(); error != nil {
		return nil, error
	}

	if error := user.Save(); error != nil {
		return nil, error
	}
	return &user, nil
}
