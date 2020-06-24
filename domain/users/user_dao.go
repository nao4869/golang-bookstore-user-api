package users

import (
	"fmt"
	"strings"

	// just for testing

	"github.com/nao4869/golang-bookstore-user-api/datasources/mysql/users_db"
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

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

// Get - user pointer in order to working on actual value in the memory
func (user *User) Get() *errors.RestError {
	// connect to mysql DB
	stmt, error := users_db.Client.Prepare(queryInsertUser)
	if error != nil {
		fmt.Println("error when trying to prepare save user statement")
		return errors.NewInternalServerError(fmt.Sprintf("error for saving user", error.Error()))
	}
	defer stmt.Close()

	// the reason passing pointer is because we want to pass a copy but not updating the actual values
	// query by user id and get single row
	result := stmt.QueryRow(user.ID)

	if getError := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); error != nil {
		if strings.Contains(getError.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to retrieve user %d", user.ID))
	}
	return nil
}

// Save - save the user to the database
func (user *User) Save() *errors.RestError {
	// insert new user to DB - creating statement connect to the DB so we must defer after communicating with it
	stmt, error := users_db.Client.Prepare(queryInsertUser)
	if error != nil {
		fmt.Println("error when trying to prepare save user statement")
		return errors.NewInternalServerError(fmt.Sprintf("error for saving user", error.Error()))
	}
	defer stmt.Close()

	// Exec return Result & Error
	insertResult, saveError := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveError != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error for saving user", error.Error()))
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		fmt.Println("error when trying to get last insert id after creating a new user")
		return errors.NewInternalServerError(fmt.Sprintf("error for saving the user to DB", err.Error()))
	}
	user.ID = userID

	return nil
}
