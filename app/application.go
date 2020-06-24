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
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.Create)
	router.Run(":8080")
}
