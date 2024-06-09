package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/auth"
	"todolist/helper"
	"todolist/user"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{
		userService,
		authService,
	}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	var err = c.ShouldBindJSON(&input)
	if err != nil {
		var errors = helper.FormatValidationError(err)
		var errorMessage = gin.H{"errors": errors}
		var _output_ = helper.APIResponse("Gagal Mengkaitkan ke JSON", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, _output_)
		return
	}
	var newUser, err2 = h.userService.RegisterUser(input)
	if err2 != nil {
		var _output_ = helper.APIResponse("Transaksi Database Gagal", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, _output_)
		return
	}
	var token, err3 = h.authService.GenerateToken(newUser.ID, newUser.Name, newUser.Password)
	if err3 != nil {
		_output_ := helper.APIResponse("Gagal Menghasilkan Token", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, _output_)
		return
	}
	var formatResponse = user.FormatUser(newUser, token)
	var _output_ = helper.APIResponse("Akun Berhasil Didaftarkan", http.StatusOK, "sukses", formatResponse)
	c.JSON(http.StatusOK, _output_)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	var err = c.ShouldBindJSON(&input)
	if err != nil {
		var errors = helper.FormatValidationError(err)
		var errorMessage = gin.H{"errors": errors}
		var _output_ = helper.APIResponse("Gagal Mengkaitkan ke JSON", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, _output_)
		return
	}
	var LoginUser, err2 = h.userService.Login(input)
	if err2 != nil {
		var errorMessage = gin.H{"errors": err.Error()}
		var _output_ = helper.APIResponse("Gagal Login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, _output_)
		return
	}
	var token, err3 = h.authService.GenerateToken(LoginUser.ID, LoginUser.Name, LoginUser.Password)
	if err3 != nil {
		var errorMessage = gin.H{"errors": err.Error()}
		var _output_ = helper.APIResponse("Gagal Login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, _output_)
		return
	}
	var formatResponse = user.FormatUser(LoginUser, token)
	var _output_ = helper.APIResponse("User Berhasil Login ", http.StatusOK, "sukses", formatResponse)
	c.JSON(http.StatusOK, _output_)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput
	var err = c.ShouldBindJSON(&input)
	if err != nil {
		var errors = helper.FormatValidationError(err)
		var errorMessage = gin.H{"errors": errors}
		var _output_ = helper.APIResponse("Gagal Mengkaitkan ke JSON", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, _output_)
		return
	}
	var isEmailAvaiable, err2 = h.userService.IsEmailAvailable(input)
	if err2 != nil {
		var errorMessage = gin.H{"errors": "Server Error"}
		var _output_ = helper.APIResponse("Pemeriksaan Email Gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, _output_)
		return
	}
	var data = gin.H{"is_available": isEmailAvaiable}
	var metaMessage = "Email Tersedia"
	if !isEmailAvaiable {
		metaMessage = "Email Sudah Digunakan"
	}
	var _output_ = helper.APIResponse(metaMessage, http.StatusOK, "sukses", data)
	c.JSON(http.StatusOK, _output_)
}
