package src

import (
	"api/middleware"
	"api/service"

	"github.com/gin-gonic/gin"
)

func AddCommentRouter(r *gin.RouterGroup) {
	comment := r.Group("/comments")
	//comment.GET("/all", service.FindAllComment)
	comment.GET("/course", service.GetCourseByQuestion)

	comment.GET("/me", middleware.RequireAuth, service.GetCommentByUserId)
	comment.POST("/me", middleware.RequireAuth, service.PostNewComment)
	comment.PATCH("/me", middleware.RequireAuth, service.PATCHCommentById)
	comment.DELETE("/me", middleware.RequireAuth, service.DeleteCommentById)

	comment.GET("/search", service.GetCommentByQuestion)
	comment.GET("/comment", service.GetCommentById)

	comment.GET("/reply", service.GetReplyById)
	comment.POST("/reply", middleware.RequireAuth, service.PostNewReply)
	comment.PATCH("/reply", middleware.RequireAuth, service.UpadteReply)
	comment.DELETE("/reply", middleware.RequireAuth, service.DeleteReplyById)

	// comment.POST("/", service.POSTAllComment)
}

func AddOauthrouter(r *gin.RouterGroup) {
	oauth := r.Group("/oauth")
	oauth.GET("/login", service.Nycu_Oauth_redirect)
	oauth.GET("/code", service.Nycu_Oauth_Get_JWT)
	oauth.GET("/me", middleware.RequireAuth, service.Nycu_check_info)
	oauth.GET("/logout", middleware.RequireAuth, service.Nycu_delete_info)

	//oauth.POST("/token", service.Nycu_Oauth_Get_token)

}
