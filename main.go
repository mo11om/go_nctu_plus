package main

import (
	"api/database"
	"api/middleware"
	. "api/src"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())
	v1 := router.Group("api/v1")
	AddCommentRouter(v1)
	AddOauthrouter(v1)

	go func() {
		database.DBconnect()
	}()
	router.Run("localhost:8080")
}
