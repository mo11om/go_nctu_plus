package controllers

import (
	"api/database"
	"fmt"
	"time"
)

type Reply struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	Content    string    `gorm:"content" json:"content" `
	Updated_at time.Time `gorm:"updated_at" json:"updated_at"`
	Created_at time.Time `gorm:"created_at" json:"created_at"`
	Name       string    `gorm:"name" json:"name"`
}

func FindreplyByDiscussId(id string) []Reply {
	var c []Reply
	query := `
	select  replys.* , users.name
	from replys
	 
	inner join users on  users .id =replys.user_id 
	where replys.discuss_id = ?
	 `
	database.Db.Raw(query,
		id).Scan(&c)

	return c
}
func CreateReply(discussId int, userId int, content string, contentType string, createdAt time.Time, updatedAt time.Time) error {
	query := `INSERT INTO replys (discuss_id, user_id, content_type, content, created_at, updated_at)
              VALUES (?, ?, ?, ?, ?, ?)`
	err := database.Db.Exec(query, discussId, userId, contentType, content, createdAt, updatedAt).Error
	return err
}
func checkUserId_is_same_to_reply_id(replyId, userId int) error {

	var user_id int
	query := `select user_id FROM replys WHERE id = ?`
	database.Db.Raw(query, replyId).Scan(&user_id)
	if user_id == userId {
		return nil

	}
	return fmt.Errorf("not same")

}
func UpdateReply(replyId int, userId int, content string) error {
	err_of_userid := checkUserId_is_same_to_reply_id(replyId, userId)
	if err_of_userid != nil {
		return err_of_userid
	}
	query := `	update  replys
				set content = ? , updated_at=?
				
				WHERE id =?     
			  `
	err := database.Db.Exec(query, content, get_time(), replyId).Error
	return err
}
func DeleteReply(replyId, user_id int) error {

	err := checkUserId_is_same_to_reply_id(replyId, user_id)
	if err != nil {
		return err
	}
	query := `DELETE FROM replys WHERE id = ?`
	err = database.Db.Exec(query, replyId).Error
	return err
}
