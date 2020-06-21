package services

import (
	"net/http"

	"github.com/nao4869/golang-bookstore-user-api/domain/users"
	"github.com/nao4869/golang-bookstore-user-api/utils/errors"
)

// CreateUser - must not return both users and errors - only return one 
// if we return nil, nil - this will cause many error in users_controllers when handling nil error and returing nil result etc
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil

	// create new instance of user struct - allocate and memery for user
	// var defaultUser users.User

	// return user, &errors.RestError{
	// 	Status: http.StatusInternalServerError,
	// }
}
