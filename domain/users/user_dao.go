package users

import (
	"fmt"

	"github.com/nao4869/golang-bookstore-user-api/utils/errors"
)

// mock user DB
var (
	usersDB = make(map[int64]*User)
)

// Get -
func (user User) Get() *errors.RestError {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID == result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	
	return nil
}

// Save - save the user to the database
func (user User) Save() *errors.RestError {
	return nil
}
