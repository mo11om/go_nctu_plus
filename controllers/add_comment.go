package controllers

import (
	"api/database"
	"fmt"
	"time"
)

type Course struct {
	Id int `json:"id"`

	Name    string `gorm:"name" json "name"`
	Ch_name string `gorm:"ch_name" json "ch_name"`
}
type NewComment struct {
	User_id               int
	Course_teachership_id int    `json:"id"`
	Title                 string `json:"title"`
	Content               string `json:"content"`
	Is_anonymous          int    `json:"is_anonymous"`
}

func get_time() string {
	theTime := time.Now()

	time_string := fmt.Sprintln(theTime.Format("2006-1-2 15:4:5"))
	return time_string
}

func FindCourseByTeacher(questions string) []Course {
	var c []Course
	var query string = "%" + questions + "%"
	database.Db.Raw("SELECT   teachers.name , courses.ch_name  ,ct.id	FROM course_teacherships as ct	 INNER JOIN courses ON courses.id = ct.course_id			  INNER JOIN teachers ON   ct.teacher_id LIKE CONCAT('[', teachers.id, ']')		 where(teachers.name like ?) 	limit 20",
		query).Scan(&c)

	return c
}
func FindCourseByID(questions string) Course {
	var c Course
	var query string = questions
	database.Db.Raw("SELECT   teachers.name , courses.ch_name  ,ct.id	FROM course_teacherships as ct	 INNER JOIN courses ON courses.id = ct.course_id			  INNER JOIN teachers ON   ct.teacher_id LIKE CONCAT('[', teachers.id, ']')		 where(ct.id	= ?) 	limit 20",
		query).Scan(&c)

	return c
}
func FindCommentByUserId(id string) []Comment {
	var c []Comment

	database.Db.Raw("SELECT  discusses.id,discusses.title,    teachers.name , courses.ch_name FROM  course_teacherships as ct INNER JOIN courses ON courses.id = ct.course_id		INNER JOIN discusses  ON       discusses .course_teachership_id = ct.id 		INNER JOIN teachers ON   ct.teacher_id LIKE CONCAT('[', teachers.id, ']')		where(discusses.user_id = ?)  ",
		id).Scan(&c)

	return c
}

func AddCommentByCourseId(newComment NewComment) error {

	user_id, course_teachership_id, title, content, is_anonymous :=
		newComment.User_id,
		newComment.Course_teachership_id,
		newComment.Title,
		newComment.Content,
		newComment.Is_anonymous
	sql_query := "INSERT INTO discusses(  user_id,course_teachership_id, title , content,is_anonymous,created_at,updated_at) VALUES ( ?,?,?,?,?,?,?);"
	time_string := get_time()

	err := database.Db.Exec(sql_query,
		user_id,
		course_teachership_id,
		title,
		content,
		is_anonymous,
		time_string,
		time_string,
	).Error
	if err != nil {
		println(err)
		return err
	}
	return nil

}
