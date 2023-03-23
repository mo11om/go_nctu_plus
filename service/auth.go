package service

import (
	"api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Nycu_Oauth_redirect(ctx *gin.Context) {

	ctx.Redirect(http.StatusTemporaryRedirect, controllers.Get_login_uri())
}

func Nycu_Oauth_Get_JWT(ctx *gin.Context) {

	code := ctx.Query("code")
	jwt_token := controllers.Get_jwt_token(code)
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", jwt_token, 3600*24, "", "", true, true)

	ctx.Redirect(http.StatusPermanentRedirect, "http://localhost:5173")

}
func Nycu_check_info(ctx *gin.Context) {
	println(ctx.GetString("user_id"))
	ctx.JSON(http.StatusOK, gin.H{
		"student_id": ctx.GetString("student_id"),
	})
}
