package pojo

import (
	"api/database"
	"time"
)

// type Page struct{

// 	PageNum int `json:page`
// }

type Comment struct {
	Id                  int       `json:"id"`
	UserId              int       `json:"user_id"`
	Courseteachershipid int       `gorm:"column:course_teachership_id"  json:"courseteachershipid"`
	Content             string    `gorm:"content" json:"Content" `
	Is_anonymous        bool      `gorm:"is_anonymous" json:"is_anonymous"`
	Title               string    `gorm:"title" json:"title"`
	Updated_at          time.Time `gorm:"updated_at" json:"updated_at"`
	Created_at          time.Time `gorm:"created_at" json:"created_at"`
}

type SendInformation struct {
	Content      string    `gorm:"content" json:"Content" `
	Is_anonymous bool      `gorm:"is_anonymous" json:"is_anonymous"`
	Title        string    `gorm:"title" json:"title"`
	Updated_at   time.Time `gorm:"updated_at" json:"updated_at"`
	Created_at   time.Time `gorm:"created_at" json:"created_at"`
}

func get_sendinformtion(c []Comment) []SendInformation {

	var send []SendInformation
	for _, c2 := range c {

		send = append(send, SendInformation{c2.Content, c2.Is_anonymous, c2.Title, c2.Created_at, c2.Updated_at})
	}
	return send
}
func FindCommentByQuestion(question string) []SendInformation {
	var c []Comment
	// var send []SendInformation
	database.Db.Raw("select * from discusses where course_teachership_id in (select id    from course_teacherships  where(  INSTR( teacher_id ,(select id from teachers where name =?)  ) > 0));",
		question).Scan(&c)

	return get_sendinformtion(c)
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
