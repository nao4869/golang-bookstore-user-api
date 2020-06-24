package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nao4869/golang-bookstore-user-api/controllers/ping"
	"github.com/nao4869/golang-bookstore-user-api/controllers/users"
)

var (
	router = gin.Default()
)

// StartApplication -
func StartApplication() {
	// url endpoints mappings
	router.GET("/ping", ping.Ping)
	//router.GET("/users/:user_id", controllers.GetUser)
	//router.GET("users/search", controllers.SearchUser)

	router.POST("/users", users.CreateUser)
	router.Run(":8080")
}
