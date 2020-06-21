package services

import "github.com/nao4869/golang-bookstore-user-api/domain/users"

// CreateUser -
func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
