package service

import (
	"api/controllers"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

	pages_query := ctx.DefaultQuery("page", "0")
	page, err := strconv.Atoi(pages_query)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "")
		return
	}
	if question == "" {
		comment, err := controllers.CommentLimitOffset(30, page)
		if err != nil {
			ctx.JSON(http.StatusNotFound, "")
			return
		}

		ctx.JSON(http.StatusOK, comment)
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
	comment := controllers.FindreplyByDiscussId(question)
	if comment == nil {
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
func get_user_id(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusNotFound, "")
		return 0, fmt.Errorf("user_id not found")
	}
	var user_id string = fmt.Sprint(id)

	tmp, err := strconv.Atoi(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0, err
	}
	return tmp, nil

}
func PostNewComment(ctx *gin.Context) {

	var newComment controllers.NewComment
	if err := ctx.ShouldBindJSON(&newComment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := get_user_id(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newComment.User_id = id
	fmt.Println(newComment)

	controllers.AddCommentByCourseId(newComment)

	// Do something with the new comment, e.g. save it to a database

	ctx.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
}

func PATCHCommentById(ctx *gin.Context) {
	var newComment controllers.NewComment
	if err := ctx.ShouldBindJSON(&newComment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := get_user_id(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newComment.User_id = id
	fmt.Println(newComment.User_id, newComment.Course_teachership_id, newComment.Is_anonymous, newComment.Title, newComment.Content)
	comment := controllers.FindCommentById(strconv.Itoa(newComment.Course_teachership_id))
	fmt.Println(comment.UserId)
	if comment.UserId != newComment.User_id {

		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	// Do something with the new comment, e.g. save it to a database
	//func PatchDiscussById(user_id, id, is_anonymous int, title, content string) error {
	err = controllers.PatchDiscussById(newComment.User_id, newComment.Course_teachership_id, newComment.Is_anonymous, newComment.Title, newComment.Content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Comment edit successfully"})

}
func PostNewReply(ctx *gin.Context) {

	var reply controllers.Reply
	if err := ctx.ShouldBindJSON(&reply); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := get_user_id(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reply.UserId = id
	//fmt.Println(reply)
	//CreateReply(discussId int, userId int, content string, contentType string, createdAt time.Time, updatedAt time.Time)
	controllers.CreateReply(reply.Id, reply.UserId, reply.Content, "1", time.Now(), time.Now())

	// Do something with the new comment, e.g. save it to a database

	ctx.JSON(http.StatusOK, gin.H{"message": "Reply created successfully"})
}
