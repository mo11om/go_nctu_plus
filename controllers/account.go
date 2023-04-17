package controllers

import (
	"api/database"
	"fmt"
	"strconv"
	"time"
)

type NCTU_User struct {
	UserId int `json:"-"`

	Name       string    `gorm:"name" json "name"`
	Student_id string    `gorm:"student_id" json "student_id"`
	Updated_at time.Time `gorm:"updated_at" json:"updated_at"`
	Created_at time.Time `gorm:"created_at" json:"created_at"`
}

func FindUserByStudent_Id(student_id string) NCTU_User {
	var user NCTU_User
	// var send []Comment

	database.Db.Raw("select * from auth_nctus where (  student_id =? )  ",
		student_id).Scan(&user)

	return user
}

func createAuthUser(userID, studentID, email string) error {
	now := get_time()

	// Execute raw SQL script
	result := database.Db.Exec("INSERT INTO auth_nctus (user_id, student_id, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		userID, studentID, email, now, now)

	// Check for errors
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(" create auth users success")
	return nil
}
func CreateUser(name, email string) error {
	query := `
        INSERT INTO users (  created_at, name, email)
        VALUES (  ?, ?, ?)
    `
	var id int

	createdAt := get_time()
	err := database.Db.Exec(query, createdAt, name, email).Error
	if err != nil {
		return err
	}
	fmt.Println("create users success")
	database.Db.Raw("select id from users where  email = ?", email).Scan(&id)
	fmt.Println(" find users success")
	userID := strconv.Itoa(id)

	createAuthUser(userID, name, email)
	return nil
}
