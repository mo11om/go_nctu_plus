package src
import (
	"github.com/gin-gonic/gin"
	"api/service"
)
func AddCommentRouter(r* gin.RouterGroup){
	comment	:= r.Group("/comments")
	comment.GET("/",service.FindAllComment)
	comment.GET("/:id",service.GetUserById)
	comment.POST("/",service.POSTAllComment)
}
