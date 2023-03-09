package controllers

import (
	"api/database"
	"time"
)

// type Page struct{

// 	PageNum int `json:page`
// }

type Comment struct {
	Id                  int       `json:"id"`
	UserId              int       `json:"-"`
	Courseteachershipid int       `gorm:"column:course_teachership_id"  json:"-"`
	Content             string    `gorm:"content" json:"-" `
	Is_anonymous        bool      `gorm:"is_anonymous" json:"is_anonymous"`
	Name                string    `gorm:"name" json "name"`
	Ch_name             string    `gorm:"ch_name" json "ch_name"`
	Title               string    `gorm:"title" json:"title"`
	Updated_at          time.Time `gorm:"updated_at" json:"updated_at"`
	Created_at          time.Time `gorm:"created_at" json:"created_at"`
}

// func FindTeacher(question string) []Teachers {

// 	var teacher []Teachers

// 	// var send []Comment

// 	database.Db.Raw("SELECT * FROM    teachers where (name =? );",
// 		question).Scan(&teacher)

// 	return teacher
// }
func FindCommentByTeacher(teacher string) []Comment {
	var c []Comment
	// var send []Comment

	database.Db.Raw("SELECT  discusses.id,discusses.title,   teachers.name , courses.ch_name FROM  course_teacherships as ct INNER JOIN courses ON courses.id = ct.course_id		INNER JOIN discusses  ON       discusses .course_teachership_id = ct.id 		INNER JOIN teachers ON   ct.teacher_id LIKE CONCAT('[', teachers.id, ']')		where(		  teachers.name =  ?		)  ",
		teacher).Scan(&c)

	return c
}

func FindCommentByChName(ch_name string) []Comment {

	var c []Comment
	var ch_name_query string = "%" + ch_name + "%"
	database.Db.Raw("SELECT  discusses.id,discusses.title,   teachers.name , courses.ch_name FROM  course_teacherships as ct INNER JOIN courses ON courses.id = ct.course_id		INNER JOIN discusses  ON       discusses .course_teachership_id = ct.id 		INNER JOIN teachers ON   ct.teacher_id LIKE CONCAT('[', teachers.id, ']')		where(courses. ch_name  like ?)  ",
		ch_name_query).Scan(&c)

	return c
}

// func FindCommentByTitle(title string) []Comment {
// 	var c []Comment
// 	var title_query string = "%" + title + "%"
// 	database.Db.Raw("SELECT  discusses.*,   teachers.name , courses.ch_name FROM  course_teacherships as ct INNER JOIN courses ON courses.id = ct.course_id		INNER JOIN discusses  ON       discusses .course_teachership_id = ct.id 		INNER JOIN teachers ON   ct.teacher_id LIKE CONCAT('[', teachers.id, ']')		where(		discusses.title like ?	) limit 20",
// 		title_query).Scan(&c)

// 	return c
// }

func FindAllCommentsByQuestion(question string) []Comment {
	// var send []Comment
	// Create channels to synchronize the goroutines
	c_teacher_ch := make(chan []Comment)
	c_chname_ch := make(chan []Comment)
	//c_title_ch := make(chan []Comment)
	go func() {
		c_chname := FindCommentByChName(question)
		c_chname_ch <- c_chname // send c_chname to the channel when the function completes
	}()
	go func() {
		c_teacher := FindCommentByTeacher(question)
		c_teacher_ch <- c_teacher // send c_teacher to the channel when the function completes
	}()

	// go func() {
	// 	c_title := FindCommentByTitle(question)
	// 	c_title_ch <- c_title // send c_chname to the channel when the function completes

	// }()

	// Wait for the two channels to receive values
	c_teacher := <-c_teacher_ch
	c_chname := <-c_chname_ch
	//c_title := <-c_title_ch
	// var c_teacher []Comment = FindCommentByTeacher(question)
	// var c_chname []Comment = FindCommentByChName(question)
	c := (append(c_chname, c_teacher...))

	//c := append(append(c_teacher, c_chname...), c_title...)
	return c
}
func FindCommentByQuestion(question string) []Comment {
	return FindAllCommentsByQuestion(question)
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
