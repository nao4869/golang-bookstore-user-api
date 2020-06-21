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
	router.GET("/ping", controllers.Ping)
	router.Run()
}
