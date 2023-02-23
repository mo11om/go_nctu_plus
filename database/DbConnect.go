package database
import (
	 
    "os"

     
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  )
  
  var Db *gorm.DB
  var Err error


  
func DBconnect() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := os.Getenv("DB")
	Db, Err   = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if Err!= nil {
         panic(Err)
    }


}