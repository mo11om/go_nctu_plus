package controllers

import (
	"api/database"
	"time"
)

type reply struct {
	Id                     int       `json:"id"`
	UserId                 int       `json:"-"`
	Course_teacher_ship_id int       `gorm:"column:course_teachership_id"  json:"-"`
	Content                string    `gorm:"content" json:"content" `
	Is_anonymous           bool      `gorm:"is_anonymous" json:"is_anonymous"`
	Updated_at             time.Time `gorm:"updated_at" json:"updated_at"`
	Created_at             time.Time `gorm:"created_at" json:"created_at"`
}

func FindreplyByDiscussId(id string) []reply {
	var c []reply
	database.Db.Raw("select * from replys  where discuss_id =?",
		id).Scan(&c)

	return c
}
