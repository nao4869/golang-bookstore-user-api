// package users

// import (
// 	"strings"

// 	"github.com/nao4869/golang-bookstore-user-api/utils/errors"
// )

// // User - use `json:"id"` value to populate struct field ID
// type User struct {
// 	ID          int64  `json:"id"`
// 	FirstName   string `json:"first_name"`
// 	LastName    string `json:"last_name"`
// 	Email       string `json:"email"`
// 	DateCreated string `json:"date_created"`
// }

// // Validate - assigning Validate method to user struct
// func (user *User) Validate() *errors.RestError {
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.NewBadRequestError("invalid email address")
// 	}
// 	return nil
// }
