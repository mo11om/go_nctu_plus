package pojo 

import (
	 "api/database"
)


type Page struct{

	PageNum int `json:page`
}

type Comment struct{
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Courseteachershipid int `gorm:"column:course_teachership_id"  json:"courseteachershipid"`
	Content string `json:content`
	Is_anonymous bool `json:"is_anonymous"`
	Title string `json:"title"`
}

func FindCommentId( id string ) Comment {
	var c Comment
	database.Db.Raw("SELECT * FROM discusses WHERE course_teachership_id=?", id).Scan(&c)
	println(c.Courseteachershipid)
	return	c	
}
// func FindCommentId( id string ) Comment {
// 	var c Comment
// 	database.Db.Raw("SELECT * FROM discusses WHERE course_teachership_id=?", id).Scan(&c)
// 	println(c.Courseteachershipid)
// 	return	c	
// }
// func FindCommentId( id string ) Comment {
// 	var c Comment
// 	database.Db.Raw("SELECT * FROM discusses WHERE course_teachership_id=?", id).Scan(&c)
// 	println(c.Courseteachershipid)
// 	return	c	
// }