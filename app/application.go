package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nao4869/golang-bookstore-user-api/controllers"
)

var (
	router = gin.Default()
)

// StartApplication -
func StartApplication() {
	// url endpoints mappings
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("users/search", users.SearchUser)

	router.POST("/users", users.CreateUser)

	router.Run()
}
