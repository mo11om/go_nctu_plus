package main

import (
	"api/database"
	"api/middleware"
	"api/src"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(middleware.Cors())
	v1 := router.Group("api/v1")
	src.AddCommentRouter(v1)
	src.AddOauthrouter(v1)

	go func() {
		database.DBconnect()
	}()
	router.Run(":8080")
}
