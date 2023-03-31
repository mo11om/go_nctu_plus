package service

import (
	"api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourseByID(ctx *gin.Context) {

	question := ctx.DefaultQuery("id", "")
	if question == "" {
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	comment := controllers.FindCourseByID(question)
	if comment.Id == 0 {
		ctx.JSON(http.StatusNotFound, "")
		return
	}

	ctx.JSON(http.StatusOK, comment)
}
func GetCourseByTeacher(ctx *gin.Context) {

	question := ctx.DefaultQuery("q", "")
	if question == "" {
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	comment := controllers.FindCourseByTeacher(question)
	if comment == nil {
		ctx.JSON(http.StatusNotFound, "")
		return
	}

	ctx.JSON(http.StatusOK, comment)
}
func GetCommentByQuestion(ctx *gin.Context) {

	question := ctx.DefaultQuery("q", "")
	if question == "" {
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	comment := controllers.FindCommentByQuestion(question)
	if comment == nil {
		ctx.JSON(http.StatusNotFound, "")
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func GetCommentById(ctx *gin.Context) {

	question := ctx.DefaultQuery("id", "")
	if question == "" {
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	comment := controllers.FindCommentById(question)
	if comment.Id == 0 {
		ctx.JSON(http.StatusNotFound, "")
		return
	}

	ctx.JSON(http.StatusOK, comment)
}
func GetReplyById(ctx *gin.Context) {

	question := ctx.DefaultQuery("id", "")
	if question == "" {
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	comment := controllers.FindreplyByCourseId(question)
	if comment == nil {
		ctx.JSON(http.StatusNotFound, "")
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// func POSTAllComment(ctx *gin.Context) {
// 	Page := controllers.Page{}
// 	err := ctx.BindJSON(&Page)
// 	if err == nil {
// 		ctx.JSON(http.StatusNotAcceptable, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, "Page post success")
// }
