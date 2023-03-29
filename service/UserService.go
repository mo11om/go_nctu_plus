package service

import (
	"api/controllers"
	"fmt"
	"net/http"
	"strconv"

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
func GetCommentByTeacher(ctx *gin.Context) {

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
func GetCommentByUserId(ctx *gin.Context) {

	//question := ctx.DefaultQuery("id", "")
	// if question == "" {
	// 	ctx.JSON(http.StatusNotFound, "")
	// 	return
	// }
	question, ok := ctx.Get("user_id")
	fmt.Println("question", question)
	if !ok {
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	id := fmt.Sprintf("%v", question)
	fmt.Println("id", id)
	comment := controllers.FindCommentByUserId(id)
	if comment == nil {
		ctx.JSON(http.StatusNotFound, "")
		return
	}

	ctx.JSON(http.StatusOK, comment)
}
func PostNewComment(ctx *gin.Context) {

	var newComment controllers.NewComment
	if err := ctx.ShouldBindJSON(&newComment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	var user_id string = fmt.Sprint(id)

	tmp, err := strconv.Atoi(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newComment.User_id = tmp

	fmt.Println(newComment)

	controllers.AddCommentByCourseId(newComment)

	// Do something with the new comment, e.g. save it to a database

	ctx.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
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
