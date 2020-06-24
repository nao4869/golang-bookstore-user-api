package users

import (
	"errors"
	"strings"

	"github.com/federicoleon/bookstore_utils-go/logger"
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
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}
	defer stmt.Close()

	// the reason passing pointer is because we want to pass a copy but not updating the actual values
	// query by user id and get single row
	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}
	return nil
}

// Save - save the user to the database
func (user *User) Save() rest_errors.RestErr {
	// insert new user to DB - creating statement connect to the DB so we must defer after communicating with it
	stmt, error := users_db.Client.Prepare(queryInsertUser)
	if error != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
		//mysql_utils.ParseError(error)
	}
	defer stmt.Close()

	// Exec return Result & Error
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		logger.Error("error when trying to save user", saveErr)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	user.ID = userID
	return nil
}