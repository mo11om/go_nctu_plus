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

func AddOauthrouter(r *gin.RouterGroup) {
	oauth := r.Group("/oauth")
    oauth.GET("/login", service.Nycu_Oauth_redirect)
    oauth.GET("/code", service.Nycu_Oauth_Get_code)
    //oauth.POST("/token", service.Nycu_Oauth_Get_token)
	 
   
}
