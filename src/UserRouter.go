package src

import (
	"api/service"

	"github.com/gin-gonic/gin"
)

func AddCommentRouter(r *gin.RouterGroup) {
	comment := r.Group("/comments")
	comment.GET("/all", service.FindAllComment)
	comment.GET("/comment", service.GetCommentById)
	comment.GET("/serach", service.GetCommentByTeacher)
	// comment.POST("/", service.POSTAllComment)
}
