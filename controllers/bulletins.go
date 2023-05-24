package controllers

import (
	"api/database"
	"time"
)

type Article struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	ArticleType int8      `gorm:"type:tinyint(1)"`
	BeginTime   time.Time `gorm:"type:datetime"`
	EndTime     time.Time `gorm:"type:datetime"`
	HiddenType  int8      `gorm:"type:tinyint(1)"`
	UserID      int       `gorm:"column:user_id"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	UpdateUser  string    `gorm:"column:update_user"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	Article     string    `gorm:"type:text"`
}

func GetNewestArticle() (Article, error) {
	var article Article
	err := database.Db.Raw("SELECT id,title FROM  bulletins ORDER BY created_at DESC LIMIT 5").Scan(&article).Error
	if err != nil {
		return article, err
	}
	return article, nil
}
