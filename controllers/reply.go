package controllers

import (
	"api/database"
	"time"
)

type Reply struct {
	Id         int       `json:"id"`
	UserId     int       `json:"-"`
	Content    string    `gorm:"content" json:"content" `
	Updated_at time.Time `gorm:"updated_at" json:"updated_at"`
	Created_at time.Time `gorm:"created_at" json:"created_at"`
}

func FindreplyByDiscussId(id string) []Reply {
	var c []Reply
	database.Db.Raw("select * from replys  where discuss_id =?",
		id).Scan(&c)

	return c
}
func CreateReply(discussId int, userId int, content string, contentType string, createdAt time.Time, updatedAt time.Time) error {
	query := `INSERT INTO replys (discuss_id, user_id, content_type, content, created_at, updated_at)
              VALUES (?, ?, ?, ?, ?, ?)`
	err := database.Db.Exec(query, discussId, userId, contentType, content, createdAt, updatedAt).Error
	return err
}
