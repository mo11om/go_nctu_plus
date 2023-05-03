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
	Content string `gorm:"content" json:"content" `
}
type NewComment struct {
	User_id               int    `json:"user_id"`
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
func FindCourseByQuestion(questions string) []Course {
	var c []Course
	var query string = "%" + questions + "%"
	sql_query := `SELECT   teachers.name , courses.ch_name  ,ct.id	
	FROM course_teacherships as ct	 
	INNER JOIN courses ON courses.id = ct.course_id			
	INNER JOIN teachers ON   ct.teacher_id = CONCAT('[', teachers.id, ']')		 
	where(teachers.name like ?) or 	(courses.ch_name  like  ?)
	limit 50
	`
	database.Db.Raw(sql_query, query, query).Scan(&c)

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

	database.Db.Raw("SELECT  discusses.id,discusses.title,  discusses.content,  teachers.name , courses.ch_name FROM  course_teacherships as ct INNER JOIN courses ON courses.id = ct.course_id		INNER JOIN discusses  ON       discusses .course_teachership_id = ct.id 		INNER JOIN teachers ON   ct.teacher_id LIKE CONCAT('[', teachers.id, ']')		where(discusses.user_id = ?)  ",
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
func CheckUserId_is_same_to_comment(userid int, id int) error {
	var userID int

	err := database.Db.Raw("SELECT user_id FROM discusses WHERE id = ?", id).Row().Scan(&userID)
	if err != nil {
		return err
	}
	if userID == userid {
		return nil
	}

	return fmt.Errorf("userID %d does not match user_id %d in the discusses table", userID, userid)

}
func PatchDiscussById(user_id, id, is_anonymous int, title, content string) error {
	// Use the raw SQL statement to update the title and content columns
	time_string := get_time()
	if err_of_userid := CheckUserId_is_same_to_comment(user_id, id); err_of_userid != nil {
		return err_of_userid
	}

	if err := database.Db.Exec("UPDATE discusses SET title = ?,is_anonymous = ?, content = ?  ,updated_at=? WHERE id = ? ", title, is_anonymous, content, time_string, id).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDiscussById(id int) error {

	// Delete the discuss by id
	if err := database.Db.Exec("DELETE FROM discusses WHERE id = ?", id).Error; err != nil {
		return err
	}

	// Also delete all replies associated with the discuss
	if err := database.Db.Exec("DELETE FROM replys WHERE discuss_id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
