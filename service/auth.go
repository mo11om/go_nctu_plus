package service

import (
	"api/controllers"
	"fmt"
	"net/http"
	"os"
	"strconv"

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

	ctx.Status(http.StatusOK)
}
func NYCU_update_info(ctx *gin.Context) {
	var user controllers.NCTU_User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	id, err := get_user_id(ctx)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	fmt.Println("upadate user id", id)
	if err := controllers.UpdateUserName(id, user.Name); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"reeason": "name is duplicate"})
		return
	}
	ctx.Status(http.StatusOK)

}
func Nycu_give_info(ctx *gin.Context) {

	println("id", ctx.GetString("user_id"))
	user_id, _ := get_user_id(ctx)
	user := controllers.FindUserById(strconv.Itoa(user_id))

	ctx.JSON(http.StatusOK, gin.H{
		"student_id": ctx.GetString("student_id"),
		"user_id":    user_id,
		"name":       user.Name,
	})
}
