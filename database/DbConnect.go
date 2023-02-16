package database
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  )
  
  var Db *gorm.DB
  var Err error


  
func DBconnect() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:secret@tcp(127.0.0.1:3306)/data?charset=utf8mb4&parseTime=True&loc=Local"
	Db, Err   = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if Err!= nil {
         panic(Err)
    }


}