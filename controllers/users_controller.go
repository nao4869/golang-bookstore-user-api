package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nao4869/golang-bookstore-user-api/domain/users"
)

// CreateUser -
func CreateUser(c *gin.Context) {
	var user users.User
	bytes, error := ioutil.ReadAll(c.Request.Body)
	fmt.Println(user)
	fmt.Println(error)
	fmt.Println(bytes)

	c.String(
		http.StatusNotImplemented,
		"Not implemented",
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
