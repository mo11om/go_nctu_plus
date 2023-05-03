package controllers

import (
	"api/database"
	"fmt"
	"math/rand"
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

func FindUserById(id string) NCTU_User {
	var user NCTU_User
	// var send []Comment

	database.Db.Raw("select * from  users where (   id =? )  ",
		id).Scan(&user)
	fmt.Println(user)
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
func ramdom_createAuthUser(n int) string {

	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)

}
func CreateUser(user_id, email string) error {
	query := `
        INSERT INTO users (  created_at,updated_at, name, email)
        VALUES (  ?, ?, ?,?)
    `
	var id int

	createdAt := get_time()
	user_name := ramdom_createAuthUser(10)
	err := database.Db.Exec(query, createdAt, createdAt, user_name, email).Error
	if err != nil {
		return err
	}
	fmt.Println("create users success")
	database.Db.Raw("select id from users where  email = ?", email).Scan(&id)
	fmt.Println(" find users success")
	userID := strconv.Itoa(id)

	createAuthUser(userID, user_id, email)
	return nil
}
func check_name_exists(name string) bool {
	var id int
	query := "SELECT id FROM users WHERE name = ? LIMIT 1"
	database.Db.Raw(query, name).Scan(&id)

	return id != 0 // record not found, name doesn't exist

}

func UpdateUserName(userId int, name string) error {
	print("upadate user id", userId, name)
	if check_name_exists(name) {
		return fmt.Errorf("name already exists")
	}
	query := `	update  users
				set name = ? 				
				WHERE id =?     
			  `
	err := database.Db.Exec(query, name, userId).Error
	return err
}
