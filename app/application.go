package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication -
func StartApplication() {
	// url endpoints mappings
	router.GET("/ping", controllers.Ping)
	router.GET("/users/:user_id", controllers.GetUser)
	//router.GET("users/search", controllers.SearchUser)

	router.POST("/users", controllers.CreateUser)

	router.Run()
}
