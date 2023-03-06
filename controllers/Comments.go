package controllers

import (
	"api/database"
	"strconv"
	"time"
)

// type Page struct{

// 	PageNum int `json:page`
// }

type Comment struct {
	Id                  int       `json:"id"`
	UserId              int       `json:"-"`
	Courseteachershipid int       `gorm:"column:course_teachership_id"  json:"-"`
	Content             string    `gorm:"content" json:"Content" `
	Is_anonymous        bool      `gorm:"is_anonymous" json:"is_anonymous"`
	Title               string    `gorm:"title" json:"title"`
	Updated_at          time.Time `gorm:"updated_at" json:"updated_at"`
	Created_at          time.Time `gorm:"created_at" json:"created_at"`
}
type Teachers struct {
	Name       string    `gorm:"name" json "name"`
	real_id    string    `gorm:"real_id" json:"real_id"`
	Is_deleted bool      `gorm:"is_deleted" json:"is_deleted"`
	Id         int       `json:"id"`
	Updated_at time.Time `gorm:"updated_at" json:"updated_at"`
	Created_at time.Time `gorm:"created_at" json:"created_at"`
}

func FindTeacher(question string) Teachers {

	var teacher Teachers

	// var send []Comment

	database.Db.Raw("SELECT * FROM    teachers where (name =? );",
		question).Scan(&teacher)

	return teacher
}

func FindCommentByQuestion(question string) []Comment {
	var c []Comment
	// var send []Comment

	teacher := FindTeacher(question)
	teacher_query := "[" + strconv.Itoa(teacher.Id) + "]"
	database.Db.Raw("select * from discusses where course_teachership_id in (select id from course_teacherships where  course_teacherships .teacher_id  =  ?);",
		teacher_query).Scan(&c)

	return c
}

func FindCommentById(id string) Comment {
	var c Comment
	database.Db.Raw("select * from discusses where id = ?",
		id).Scan(&c)

	return c
}
func FindAllComment() []Comment {
	var c []Comment
	database.Db.Raw("select * from discusses ;").Scan(&c)

	return c
}
