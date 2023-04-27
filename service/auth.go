package service

import (
	"api/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Nycu_Oauth_redirect(ctx *gin.Context) {

	ctx.Redirect(http.StatusTemporaryRedirect, controllers.Get_login_uri())
}

func Nycu_Oauth_Get_JWT(ctx *gin.Context) {
	front_end_uri := os.Getenv("front_end_uri")
	fmt.Println(front_end_uri)
	code := ctx.Query("code")

	jwt_token, err := controllers.Get_jwt_token(code)
	if err != nil {
		ctx.Redirect(http.StatusUnauthorized, front_end_uri)
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", jwt_token, 3600*24, "/", "", false, true)

	ctx.Redirect(http.StatusPermanentRedirect, front_end_uri)

}
func Nycu_delete_info(ctx *gin.Context) {

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"logout": true,
	})
}
func Nycu_check_info(ctx *gin.Context) {
	println(ctx.GetString("user_id"))
	ctx.JSON(http.StatusOK, gin.H{
		"student_id": ctx.GetString("student_id"),
	})
}
