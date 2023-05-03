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
	error_of_oauth := ctx.DefaultQuery("error", "")
	fmt.Println("error_of_oauth", error_of_oauth)
	if error_of_oauth != "" {
		fmt.Println("error")
		ctx.Redirect(http.StatusUnauthorized, front_end_uri)
		return
	}

	jwt_token, err := controllers.Get_jwt_token(code)

	if err != nil {
		ctx.Redirect(http.StatusUnauthorized, front_end_uri)
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", jwt_token, 3600*24, "/", "", false, true)
	//ctx.SetCookie("Login", "true", 3600*24, "/", "", false, true)

	ctx.Redirect(http.StatusPermanentRedirect, front_end_uri)

}
func Nycu_delete_info(ctx *gin.Context) {

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)
	ctx.Status(http.StatusOK)
}
func Nycu_check_info(ctx *gin.Context) {
	println("id", ctx.GetString("user_id"))
	user_id, _ := get_user_id(ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"student_id": ctx.GetString("student_id"),
		"user_id":    user_id,
	})
}
