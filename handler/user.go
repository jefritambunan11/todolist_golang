package handler

import(
	"bwastartup/user"	
	"bwastartup/helper"	

	"github.com/gin-gonic/gin"
	
	"net/http"
	"fmt"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput 

	err := c.ShouldBindJSON(&input)
	if err != nil {				
		errors := helper.FormatValidationError(err)	
		errorMessage := gin.H{ "errors": errors }
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

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput 

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)	
		errorMessage := gin.H{ "errors": errors }
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	} 

	LoginUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{ "errors": err.Error() }
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response) 
		return	 	
	}

	formatResponse := user.FormatUser(LoginUser, "tokentokentokentokentoken")
	response := helper.APIResponse("User Successfully Loggedin", http.StatusOK, "success", formatResponse)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {	
	var input user.CheckEmailInput 

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)	
		errorMessage := gin.H{ "errors": errors }
		response := helper.APIResponse("Email Checking Is Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	} 

	isEmailAvaiable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{ "errors": "Server Error" }
		response := helper.APIResponse("Email Checking Is Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response) 
		return	 	
	}

	data := gin.H{ "is_available": isEmailAvaiable }

	var metaMessage string = "Email Is Available"
	if !isEmailAvaiable { metaMessage = "Email Has Been Taken" }

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}


func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{ "is_uploaded": false }
		response := helper.APIResponse("Failed To Upload Avatar Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 1
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	//path := "images/" + file.Filename

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{ "is_uploaded": false }
		response := helper.APIResponse("Failed To Upload Avatar Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{ "is_uploaded": false }
		response := helper.APIResponse("Failed To Upload Avatar Image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{ "is_uploaded": true }
	response := helper.APIResponse("Avatar Successfuly Uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}