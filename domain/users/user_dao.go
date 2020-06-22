package users

import (
	"fmt"
	"strings"

	// just for testing
	_ "github.com/go-sql-driver/mysql"
	usersdb "github.com/nao4869/golang-bookstore-user-api/datasources/mysql/users_db"
	date "github.com/nao4869/golang-bookstore-user-api/utils/date_utils"
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
func (user *User) Save() *errors.RestError {
	// insert new user to DB - creating statement connect to the DB so we must defer after communicating with it
	statement, error := usersdb.Client.Prepare(queryInsertUser)
	fmt.Println(statement)
	if error != nil {
		fmt.Println("Internal Server Error")
		return errors.NewInternalServerError(error.Error())
	}
	defer statement.Close()
	user.DateCreated = date.GetCurrentTimeString()

	// Exec return Result & Error
	insertResult, error := statement.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
	)
	if error != nil {
		if strings.Contains(error.Error(), "email_UNIQUE") {
			return errors.NewBadRequestError(fmt.Sprintf("email %s is already exists", user.Email))
		}
		fmt.Println("Internal Server Error - 2")
		return errors.NewInternalServerError(fmt.Sprintf("error for saving the user: %s", error.Error()))
	}
	userID, error := insertResult.LastInsertId()
	if error != nil {
		fmt.Println("Internal Server Error - 3")
		return errors.NewInternalServerError(fmt.Sprintf("error for saving the user: %s", error.Error()))
	}
	user.ID = userID // assigning last insert user id to User.ID
	//user.DateCreated = date.GetCurrentTimeString()
	return nil
}
