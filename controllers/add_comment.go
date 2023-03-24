package controllers

import (
	"api/database"
)

type Course struct {
	Id int `json:"id"`

	Name    string `gorm:"name" json "name"`
	Ch_name string `gorm:"ch_name" json "ch_name"`
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

func AddCommentByCommentId(id string) []Comment {
	var c []Comment

	database.Db.Raw("SELECT  discusses.id,discusses.title,    teachers.name , courses.ch_name FROM  course_teacherships as ct INNER JOIN courses ON courses.id = ct.course_id		INNER JOIN discusses  ON       discusses .course_teachership_id = ct.id 		INNER JOIN teachers ON   ct.teacher_id LIKE CONCAT('[', teachers.id, ']')		where(discusses.user_id = ?)  ",
		id).Scan(&c)

	return c
}
