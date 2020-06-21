package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nao4869/golang-bookstore-user-api/domain/users"
)

// CreateUser -
func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, error := ioutil.ReadAll(c.Request.Body)
	fmt.Println(bytes)
	fmt.Println(error)

	if error != nil {
		return
	}

	// trying to use given bytes json to populate user struct
	if error = json.Unmarshal(bytes, &user); error != nil {
		// TODO handle JSON error
		fmt.Println(error.Error())
		return
	}
	fmt.Println(user)

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
