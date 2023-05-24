package controllers

import (
	"api/database"
	"fmt"
	"time"
)

// type Page struct{

// 	PageNum int `json:page`
// }

type Comment struct {
	Id                  int    `json:"id"`
	User_Id             int    `gorm:"user_id"json:"user_id"`
	Courseteachershipid int    `gorm:"column:course_teachership_id"  json:"-"`
	Content             string `gorm:"content" json:"content" `
	Is_anonymous        int    `gorm:"is_anonymous" json:"is_anonymous"`

	Title      string    `gorm:"title" json:"title"`
	Updated_at time.Time `gorm:"updated_at" json:"-"`
	Created_at time.Time `gorm:"created_at" json:"-"`
	Name       string    `gorm:"name" json "name"`
	Ch_name    string    `gorm:"ch_name" json "ch_name"`
}

func FindCommentByTeacher(teacher string) []Comment {
	var c []Comment
	// var send []Comment
	query := `
	SELECT  discusses.id,discusses.title,   teachers.name , courses.ch_name 
	FROM  course_teacherships as ct 
	INNER JOIN courses ON courses.id = ct.course_id	
	INNER JOIN discusses  ON       discusses .course_teachership_id = ct.id 		
	INNER JOIN teachers ON   ct.teacher_id = CONCAT('[', teachers.id, ']')		
	where(		  teachers.name =  ?		) `
	database.Db.Raw(query,
		teacher).Scan(&c)

	return c
}

func FindCommentByChName(ch_name string) []Comment {

	var c []Comment
	var ch_name_query string = "%" + ch_name + "%"
	query := `SELECT  discusses.id,discusses.title,   teachers.name , courses.ch_name FROM  course_teacherships as ct 
	INNER JOIN courses ON courses.id = ct.course_id		
	INNER JOIN discusses  ON       discusses .course_teachership_id = ct.id 		
	INNER JOIN teachers ON   ct.teacher_id = CONCAT('[', teachers.id, ']')		
	where(courses. ch_name  like ?) `
	database.Db.Raw(query,
		ch_name_query).Scan(&c)

	return c
}

func FindCommentByTitle(title string) []Comment {
	var c []Comment
	var title_query string = "%" + title + "%"
	query := `SELECT  discusses.*,   teachers.name , courses.ch_name FROM  course_teacherships as ct 
	 INNER JOIN courses ON courses.id = ct.course_id		
	 INNER JOIN discusses  ON       discusses .course_teachership_id = ct.id 		
	 INNER JOIN teachers ON   ct.teacher_id = CONCAT('[', teachers.id, ']')		
	 where(		discusses.title like ?	) `
	database.Db.Raw(query,
		title_query).Scan(&c)

	return c
}

func FindAllCommentsByQuestion(question string) []Comment {
	// var send []Comment
	// Create channels to synchronize the goroutines
	c_teacher_ch := make(chan []Comment)
	c_chname_ch := make(chan []Comment)
	c_title_ch := make(chan []Comment)
	go func() {
		c_chname := FindCommentByChName(question)
		c_chname_ch <- c_chname // send c_chname to the channel when the function completes
	}()
	go func() {
		c_teacher := FindCommentByTeacher(question)
		c_teacher_ch <- c_teacher // send c_teacher to the channel when the function completes
	}()

	go func() {
		c_title := FindCommentByTitle(question)
		c_title_ch <- c_title // send c_chname to the channel when the function completes

	}()

	// Wait for the two channels to receive values
	c_teacher := <-c_teacher_ch
	c_chname := <-c_chname_ch
	c_title := <-c_title_ch

	//c := (append(c_chname, c_teacher...))

	c := append(append(c_teacher, c_chname...), c_title...)
	return c
}
func FindCommentByQuestion(question string) []Comment {
	//return FindAllCommentsByQuestion(question)
	var c []Comment
	var title_query string = "%" + question + "%"
	fmt.Println(title_query)
	query :=
		` 
		SELECT  discusses.id,discusses.title,   teachers.name , courses.ch_name FROM  course_teacherships as ct 
		INNER JOIN courses ON courses.id = ct.course_id		
		INNER JOIN discusses  ON       discusses .course_teachership_id = ct.id 		
		INNER JOIN teachers ON   ct.teacher_id = CONCAT('[', teachers.id, ']')		
		where(courses. ch_name  like ? )
		or  ( teachers.name like ?  ) 
		or discusses.title like ? 
		

	`
	database.Db.Raw(query,
		title_query, title_query, title_query).Scan(&c)

	return c
}

func FindCommentById(id string) Comment {
	var c Comment
	query := `SELECT  discusses.*  ,teachers.name , courses.ch_name FROM  course_teacherships as ct 
	INNER JOIN courses ON courses.id = ct.course_id		
	INNER JOIN discusses  ON       discusses .course_teachership_id = ct.id 		
	INNER JOIN teachers ON   ct.teacher_id LIKE CONCAT('[', teachers.id, ']')	
	where(discusses.id = ?)  `
	database.Db.Raw(query,
		id).Scan(&c)

	return c
}

func CommentLimitOffset(limit, page int) ([]Comment, error) {
	var c []Comment
	query := `
	SELECT discusses.id, discusses.title , teachers.name, courses.ch_name
	FROM course_teacherships as ct
	INNER JOIN courses ON courses.id = ct.course_id
	INNER JOIN discusses ON discusses.course_teachership_id = ct.id
	INNER JOIN teachers ON ct.teacher_id = CONCAT('[', teachers.id, ']')
	
	order by discusses.id desc
	LIMIT ? OFFSET ?
`
	err := database.Db.Raw(query, limit, page*limit).Scan(&c).Error
	if err != nil {
		return nil, err
	}
	return c, nil
}
