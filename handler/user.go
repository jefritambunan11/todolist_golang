package handler

import(
	"bwastartup/user"	
	"bwastartup/helper"	

	"github.com/gin-gonic/gin"
	
	"net/http"
	// "fmt"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput 
	// struct di atas kita passing sebagai parameter service

	var input user.RegisterUserInput 

	err := c.ShouldBindJSON(&input)
	if err != nil {		
		
		errors := helper.FormatValidationError(err)
	
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response) 
		return	 
	}

	formatResponse := user.FormatUser(newUser, "tokentokentokentokentoken")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatResponse)

	c.JSON(http.StatusOK, response)

}