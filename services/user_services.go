package services

import (
	"github.com/nao4869/bookstore_users-api/utils/crypto_utils"
	"github.com/nao4869/bookstore_users-api/utils/date_utils"
	"github.com/nao4869/bookstore_utils-go/rest_errors"
	"github.com/nao4869/golang-bookstore-user-api/domain/users"
)

var (
	// UsersService -
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, rest_errors.RestErr)
}

// GetUser -
// func GetUser(userID int64) (*users.User, rest_errors.RestErr) {
// 	result := &users.User{ID: userID}
// 	if error := result.Get(); error != nil {
// 		return nil, error
// 	}
// 	return result, nil
// }

// CreateUser - must not return both users and errors - only return one
// if we return nil, nil - this will cause many error in users_controllers when handling nil error and returing nil result etc
func (s *usersService) CreateUser(user users.User) (*users.User, rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
