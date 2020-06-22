package users

import (
	"fmt"
	"log"
	"strings"

	// just for testing
	_ "github.com/go-sql-driver/mysql"
	usersdb "github.com/nao4869/golang-bookstore-user-api/datasources/mysql/users_db"
	"github.com/nao4869/golang-bookstore-user-api/utils/errors"
)

// User - use `json:"id"` value to populate struct field ID
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate - assigning Validate method to user struct
func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}

// mock user DB
var (
	usersDB = make(map[int64]*User)
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
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
func (user *User) Save() error {
	// insert new user to DB - creating statement connect to the DB so we must defer after communicating with it
	statement, error := usersdb.Client.Prepare(queryInsertUser)
	if error != nil {
		return error
	}
	defer statement.Close()

	// Exec return Result & Error
	insertResult, error := statement.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
	)
	if error != nil {
		if strings.Contains(error.Error(), "email_UNIQUE") {
			return error
		}
		return error
	}
	userID, error := insertResult.LastInsertId()
	if error != nil {
		return error
	}
	user.ID = userID // assigning last insert user id to User.ID
	log.Println("Save user data")
	return nil
}
