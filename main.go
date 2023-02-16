package main

import (
	"github.com/gin-gonic/gin"
	. "api/src"
	"api/database"
 
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	AddCommentRouter(v1)
	go func() {
		database.DBconnect()
	}()
	router.Run("localhost:8080")
}
