package main

import(
	"bwastartup/user"
	"bwastartup/handler"
	
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	
	// "fmt"
	"log"
	// "net/http"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal(err.Error())
	}
	
	userRepository := user.NewRepository(db) 
	userService := user.NewService(userRepository) 
	userHandler := handler.NewUserHandler(userService)	
	
	router := gin.Default()
	api_v1 := router.Group("/api/v1", userHandler.RegisterUser)
	
	api_v1.POST("/users")
	
	router.Run(":5000")
	
}

