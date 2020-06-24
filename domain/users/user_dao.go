package users

import (
	"errors"
	"fmt"

	"github.com/nao4869/golang-bookstore-user-api/datasources/mysql/users_db"
	"github.com/nao4869/golang-bookstore-user-api/utils/rest_errors"
)

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
		return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}
	defer stmt.Close()

	// the reason passing pointer is because we want to pass a copy but not updating the actual values
	// query by user id and get single row
	result := stmt.QueryRow(user.ID)

	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}
	return nil
}

// Save - save the user to the database
func (user *User) Save() rest_errors.RestErr {
	// insert new user to DB - creating statement connect to the DB so we must defer after communicating with it
	stmt, error := users_db.Client.Prepare(queryInsertUser)
	if error != nil {
		fmt.Println(error)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	defer stmt.Close()

	// Exec return Result & Error
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		fmt.Println(saveErr)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	user.ID = userID
	return nil
}
