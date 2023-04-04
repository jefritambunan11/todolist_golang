package handler

import (
	"todolist/auth"
	"todolist/helper"
	"todolist/todo"
	"todolist/user"

	"github.com/gin-gonic/gin"

	
	"net/http"
	"strconv"
)

type todoHandler struct {
	todoService todo.Service
	authService auth.Service
}

func NewTodoHandler(todoService todo.Service, authService auth.Service) *todoHandler {
	return &todoHandler{
		todoService,
		authService,
	}
}

func (h *todoHandler) GetTodos(c *gin.Context) {
	
	userID, _ := strconv.Atoi(c.Query("user_id"))
	
	
	todos, err := h.todoService.GetTodos(userID)
	if err != nil {
		response := helper.APIResponse("Gagal Mengambil Data Todos", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	
	response := helper.APIResponse("Daftar Todos", http.StatusOK, "sukses", todo.FormatTodos(todos))
	c.JSON(http.StatusOK, response)

}

func (h *todoHandler) GetTodo(c *gin.Context) {
	
	var input todo.GetTodoDetailInput
	
	var err = c.ShouldBindUri(&input)
	if err != nil {
		var _output_ = helper.APIResponse("Gagal Mengkaitkan ke JSON ", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, _output_)
		return
	}
	
	
	var todoDetail, err2 = h.todoService.GetTodoByID(input)
	if err2 != nil {
		var _output_ = helper.APIResponse("Transaksi Database Gagal", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, _output_)
		return
	}

	var _output_ = helper.APIResponse("Todo detail", http.StatusOK, "sukses", todo.FormatTodoDetail(todoDetail))
	c.JSON(http.StatusOK, _output_)
}

func (h *todoHandler) CreateTodo(c *gin.Context) {
	
	var input todo.CreateTodoInput

	var err = c.ShouldBindJSON(&input)
	if err != nil {
		var errors = helper.FormatValidationError(err)
		var errorMessage = gin.H{"error": errors}

		var _output_ = helper.APIResponse("Transaksi Database Gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, _output_)
		return
	}

	var currentUser = c.MustGet("userAktif").(user.User)
	input.User = currentUser

	var newTodo, err2 = h.todoService.CreateTodo(input)

	if err2 != nil {
		var _output_ = helper.APIResponse("Transaksi Database Gagal", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, _output_)
		return
	}

	var _output_ = helper.APIResponse("Berhasil Membuat Todo", http.StatusOK, "sukses", todo.FormatTodoDetail(newTodo))
	c.JSON(http.StatusOK, _output_)

}

func (h *todoHandler) UpdateTodo(c *gin.Context) {
	
	var inputID todo.GetTodoDetailInput

	var err = c.ShouldBindUri(&inputID)
	if err != nil {
		var _output_ = helper.APIResponse("Gagal Mengkaitkan ke URI", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, _output_)
		return
	}

	var inputData todo.CreateTodoInput
	var err2 = c.ShouldBindJSON(&inputData)
	if err2 != nil {
		var errors = helper.FormatValidationError(err)
		var errorMessage = gin.H{"error": errors}

		var _output_ = helper.APIResponse("Gagal Mengkaitkan Ke JSON", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, _output_)
		return
	}

	var currentUser = c.MustGet("userAktif").(user.User)
	inputData.User = currentUser

	var updateTodo, err3 = h.todoService.UpdateTodo(inputID, inputData)
	if err3 != nil {
		var _output_ = helper.APIResponse("Gagal Memperbaharui Todo", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, _output_)
		return
	}

	var _output_ = helper.APIResponse("Berhasil Memperbaharui Todo", http.StatusOK, "sukses", todo.FormatTodo(updateTodo))
	c.JSON(http.StatusOK, _output_)
}

func (h *todoHandler) DeleteTodo(c *gin.Context) {

	var inputID todo.GetTodoDetailInput
	
	var err = c.ShouldBindUri(&inputID)
	if err != nil {
		var _output_ = helper.APIResponse("Gagak Mengkaitkan Ke URI", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, _output_)
		return
	}

	var inputData todo.CreateTodoInput

	var currentUser = c.MustGet("userAktif").(user.User)
	inputData.User = currentUser

	var deleteTodo, err2 = h.todoService.DeleteTodo(inputID, inputData)
	if err2 != nil {
		var _output_ = helper.APIResponse("Transaksi Database Gagal", http.StatusUnprocessableEntity, "error", err2)
		c.JSON(http.StatusUnprocessableEntity, _output_)
		return
	}

	var _output_ = helper.APIResponse("Berhasil Menghapus Todo", http.StatusOK, "sukses", todo.FormatTodo(deleteTodo))
	c.JSON(http.StatusOK, _output_)
}
