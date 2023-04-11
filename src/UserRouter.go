package src

import (
	"api/middleware"
	"api/service"

	"github.com/gin-gonic/gin"
)

func AddCommentRouter(r *gin.RouterGroup) {
	comment := r.Group("/comments")
	//comment.GET("/all", service.FindAllComment)
	comment.GET("/course", service.GetCourseByTeacher)
	comment.GET("/comment", service.GetCommentById)
	comment.GET("/add", service.GetCourseByID)
	comment.POST("/add", middleware.RequireAuth, service.PostNewComment)
	comment.GET("/search", service.GetCommentByTeacher)
	comment.GET("/me", middleware.RequireAuth, service.GetCommentByUserId)

	comment.GET("/search", service.GetCommentByQuestion)

	comment.GET("/reply", service.GetReplyById)
	// comment.POST("/", service.POSTAllComment)
}

func AddOauthrouter(r *gin.RouterGroup) {
	oauth := r.Group("/oauth")
	oauth.GET("/login", service.Nycu_Oauth_redirect)
	oauth.GET("/code", service.Nycu_Oauth_Get_JWT)
	oauth.GET("/me", middleware.RequireAuth, service.Nycu_check_info)
	//oauth.POST("/token", service.Nycu_Oauth_Get_token)

}
