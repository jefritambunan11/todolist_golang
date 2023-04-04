package main

import (
	"net/http"
	"strings"
	"todolist/auth"
	"todolist/database"
	"todolist/handler"
	"todolist/helper"
	"todolist/todo"
	"todolist/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main()  {
	
	var db = database.Connect()

	
	var router = gin.Default()
	router.SetTrustedProxies(nil)

	
	var authService = auth.NewService()

	
	var userRepository = user.NewRepository(db)
	var userService = user.NewService(userRepository)
	var userHandler = handler.NewUserHandler(userService, authService)

	
	var api = router.Group("/api")
	

	
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)

	
	var todoRepository = todo.NewRepository(db)
	var todoService = todo.NewService(todoRepository)
	var todoHandler = handler.NewTodoHandler(todoService, authService)

	
	api.GET("/todo_list", todoHandler.GetTodos)
	api.GET("/todo_list/:id", todoHandler.GetTodo)
	api.POST("/todo", authMiddleware(authService, userService), todoHandler.CreateTodo)
	api.PUT("/todo/:id", authMiddleware(authService, userService), todoHandler.UpdateTodo)
	api.DELETE("/todo/:id", authMiddleware(authService, userService), todoHandler.DeleteTodo)


	router.Run(":8080")

}


func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		var authHeader = c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			var _output_ = helper.APIResponse("Bentuk Token Salah ", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, _output_)
			return
		}

		var tokenString = ""

		var arrayToken = strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		var _token_, err = authService.ValidateToken(tokenString)
		if err != nil {
			var _output_ = helper.APIResponse("Token Tidak Valid", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, _output_)
			return
		}

		var claim, ok = _token_.Claims.(jwt.MapClaims)

		if !ok || !_token_.Valid {
			var _output_ = helper.APIResponse("Token Tidak Valid", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, _output_)
			return
		}

		var userID = int(claim["user_id"].(float64))

		var user, err2 = userService.GetUserByID(userID)
		if err2 != nil {
			var _output_ = helper.APIResponse("ID User Tidak Ditemukan", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, _output_)
			return
		}

		c.Set("userAktif", user)
	}

}
