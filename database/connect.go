package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	
	var db_config = "root:@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	
	var db, err = gorm.Open(mysql.Open(db_config), &gorm.Config{})
	if err != nil { panic(err.Error()) }
	
	return db

}
