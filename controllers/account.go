package controllers

import (
	"api/database"
	"time"
)

type NCTU_User struct {
	UserId int `json:"-"`

	Name       string    `gorm:"name" json "name"`
	Student_id string    `gorm:"student_id" json "student_id"`
	Updated_at time.Time `gorm:"updated_at" json:"updated_at"`
	Created_at time.Time `gorm:"created_at" json:"created_at"`
}

func findUserByStudent_Id(student_id string) NCTU_User {
	var user NCTU_User
	// var send []Comment

	database.Db.Raw("select * from auth_nctus where (  student_id =? )  ",
		student_id).Scan(&user)

	return user
}
