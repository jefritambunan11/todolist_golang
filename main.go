package main

import (
	"log"
	"net/http"
	"strings"
	"todolist/auth"
	"todolist/handler"
	"todolist/helper"
	"todolist/todo"
	"todolist/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)

	authService := auth.NewService()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	api := router.Group("/api")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)

	todoRepository := todo.NewRepository(db)
	todoService := todo.NewService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService, authService)

	api.GET("/todo_list", todoHandler.GetTodos)
	api.GET("/todo_list/:id", todoHandler.GetTodo)
	api.POST("/todo", authMiddleware(authService, userService), todoHandler.CreateTodo)
	api.PUT("/todo/:id", authMiddleware(authService, userService), todoHandler.UpdateTodo)
	api.DELETE("/todo/:id", authMiddleware(authService, userService), todoHandler.DeleteTodo)

	router.Run(":8080")

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Bentuk Token Salah ", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tokenString string

		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		_token_, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Token Tidak Valid", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := _token_.Claims.(jwt.MapClaims)

		if !ok || !_token_.Valid {
			response := helper.APIResponse("Token Tidak Valid", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("ID User Tidak Ditemukan", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("userAktif", user)
	}

}
