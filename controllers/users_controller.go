package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nao4869/golang-bookstore-user-api/domain/users"
	"github.com/nao4869/golang-bookstore-user-api/services"
	"github.com/nao4869/golang-bookstore-user-api/utils/errors"
)

// CreateUser -
func CreateUser(c *gin.Context) {
	var user users.User
	if error := c.ShouldBindJSON(&user); error != nil {
		//fmt.Println(error)
		// TODO: return bad request to the console
		restError := errors.RestError{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restError.Status, restError)
		return
	}

	// sending user to services CreateUser function
	result, saveError := services.CreateUser(user)
	if saveError != nil {
		// TODO handle user creating error
		return
	}
	fmt.Println(user)

	c.JSON(
		http.StatusCreated,
		result,
	)
}

// GetUser -
func GetUser(c *gin.Context) {
	c.String(
		http.StatusNotImplemented,
		"Not implemented",
	)
}

// SearchUser -
func SearchUser(c *gin.Context) {
	c.String(
		http.StatusNotImplemented,
		"Not implemented",
	)
}
