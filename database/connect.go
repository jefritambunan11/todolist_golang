package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
<<<<<<< HEAD
	// TODO 
	// 1. konfig untuk konek ke database
	// 2. jalankan konfig dan akses database dan store koneksi ke var db
	// 3. jika terjadi error keluarkan panic
	// 4. return var db
	
	// TODO 1 
	var db_config = "root:@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	
	// TODO 2
	var db, err = gorm.Open(mysql.Open(db_config), &gorm.Config{})

	// TODO 3
	if err != nil {
		panic(err.Error())
	}

	// TODO 4
=======
	
	var db_config = "root:@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	
	var db, err = gorm.Open(mysql.Open(db_config), &gorm.Config{})
	if err != nil { panic(err.Error()) }
	
>>>>>>> 9ede165dd324e1863802b8cdb43c54dc29b7457e
	return db

}
