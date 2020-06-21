package users

import (
	"fmt"

	date "github.com/nao4869/golang-bookstore-user-api/utils/date_utils"
	"github.com/nao4869/golang-bookstore-user-api/utils/errors"
)

// mock user DB
var (
	usersDB = make(map[int64]*User)
)

// Get - user pointer in order to working on actual value in the memory
func (user *User) Get() *errors.RestError {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	fmt.Println(user.ID)
	fmt.Println(user.FirstName)
	fmt.Println(user.LastName)
	fmt.Println(user.Email)
	fmt.Println(user.DateCreated)

	return nil
}

// Save - save the user to the database
func (user *User) Save() *errors.RestError {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("this email %s is already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists in DB", user.ID))
	}

	user.DateCreated = date.GetCurrentTimeString()

	usersDB[user.ID] = user
	return nil
}
