package service

import (
	"api/pojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

//get all
func FindAllComment(ctx *gin.Context) {
	commentList := pojo.FindAllComment()

	ctx.JSON(http.StatusOK, commentList)

}

func GetCommentByTeacher(ctx *gin.Context) {

	question := ctx.DefaultQuery("q", "")
	if question == "" {
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	comment := pojo.FindCommentByQuestion(question)
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
	comment := pojo.FindCommentById(question)
	if comment.Id == 0 {
		ctx.JSON(http.StatusNotFound, "")
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// func POSTAllComment(ctx *gin.Context) {
// 	Page := pojo.Page{}
// 	err := ctx.BindJSON(&Page)
// 	if err == nil {
// 		ctx.JSON(http.StatusNotAcceptable, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, "Page post success")
// }
