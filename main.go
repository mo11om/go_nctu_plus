package main

import (
	"api/database"
	. "api/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("api/v1")
	AddCommentRouter(v1)
	go func() {
		database.DBconnect()
	}()
	router.Run("localhost:8080")
}
