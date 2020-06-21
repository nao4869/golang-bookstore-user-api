package users

import (
	"fmt"

	// just for testing
	_ "github.com/go-sql-driver/mysql"
	usersdb "github.com/nao4869/golang-bookstore-user-api/datasources/mysql/users_db"
	date "github.com/nao4869/golang-bookstore-user-api/utils/date_utils"
	"github.com/nao4869/golang-bookstore-user-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
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
	// insert new user to DB - creating statement connect to the DB so we must defer after communicating with it
	statement, error := usersdb.Client.Prepare(queryInsertUser)
	if error != nil {
		return errors.NewInternalServerError(error.Error())
	}
	defer statement.Close()

	// result, error := users_db.Client.Exec(
	// 	queryInsertUser,
	// 	user.FirstName,
	// 	user.LastName,
	// 	user.Email,
	// 	user.DateCreated,
	// )

	// Exec return Result & Error
	insertResult, error := statement.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
	)
	if error != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error for saving the user: %s", error.Error()))
	}
	userID, error := insertResult.LastInsertId()
	if error != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error for saving the user: %s", error.Error()))
	}
	user.ID = userID // assigning last insert user id to User.ID
	user.DateCreated = date.GetCurrentTimeString()
	return nil
}
