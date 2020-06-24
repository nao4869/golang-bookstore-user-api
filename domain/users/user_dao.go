package users

import (
	"strings"

	// just for testing

	"github.com/federicoleon/bookstore_users-api/utils/mysql_utils"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"github.com/nao4869/golang-bookstore-user-api/datasources/mysql/users_db"
)

// User - use `json:"id"` value to populate struct field ID
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

// Validate - assigning Validate method to user struct
func (user *User) Validate() rest_errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return rest_errors.NewBadRequestError("invalid password")
	}
	return nil
}

const (
	errorNoRows     = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

// Get - user pointer in order to working on actual value in the memory
func (user *User) Get() rest_errors.RestErr {
	// connect to mysql DB
	stmt, error := users_db.Client.Prepare(queryInsertUser)
	if error != nil {
		mysql_utils.ParseError(error)
	}
	defer stmt.Close()

	// the reason passing pointer is because we want to pass a copy but not updating the actual values
	// query by user id and get single row
	result := stmt.QueryRow(user.ID)

	if getError := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); error != nil {
		return mysql_utils.ParseError(getError)
	}
	return nil
}

// Save - save the user to the database
func (user *User) Save() rest_errors.RestErr {
	// insert new user to DB - creating statement connect to the DB so we must defer after communicating with it
	stmt, error := users_db.Client.Prepare(queryInsertUser)
	if error != nil {
		mysql_utils.ParseError(error)
	}
	defer stmt.Close()

	// Exec return Result & Error
	insertResult, saveError := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveError != nil {
		return mysql_utils.ParseError(saveError)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.ID = userID
	return nil
}
