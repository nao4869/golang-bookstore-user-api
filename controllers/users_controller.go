package controllers

import (
	"net/http"

	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
	"github.com/nao4869/golang-bookstore-user-api/domain/users"
	services "github.com/nao4869/golang-bookstore-user-api/services"
)

// CreateUser -
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetUser -
// func GetUser(c *gin.Context) {
// 	// error handling for user_id
// 	userID, userError := strconv.ParseInt(c.Param("user_id"), 10, 64)
// 	if userError != nil {
// 		error := errors.NewBadRequestError("user id should be a number")
// 		c.JSON(error.Status, error)
// 		return
// 	}

// 	user, getError := services.GetUser(userID)
// 	if getError != nil {
// 		c.JSON(getError.Status, user)
// 		return
// 	}

// 	c.JSON(
// 		http.StatusOK,
// 		user,
// 	)
// }

// SearchUser -
func SearchUser(c *gin.Context) {
	c.String(
		http.StatusNotImplemented,
		"Not implemented",
	)
}
