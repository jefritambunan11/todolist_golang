package handler

import (
	"todolist/auth"
	"todolist/helper"
	"todolist/todo"
	"todolist/user"

	"github.com/gin-gonic/gin"

	// "fmt"
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

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Gagal Mengkaitkan ke JSON ", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	todoDetail, err := h.todoService.GetTodoByID(input)
	if err != nil {
		response := helper.APIResponse("Transaksi Database Gagal", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Todo detail", http.StatusOK, "sukses", todo.FormatTodoDetail(todoDetail))
	c.JSON(http.StatusOK, response)
}

func (h *todoHandler) CreateTodo(c *gin.Context) {
	var input todo.CreateTodoInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Transaksi Database Gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("userAktif").(user.User)
	input.User = currentUser

	newTodo, err := h.todoService.CreateTodo(input)

	if err != nil {
		response := helper.APIResponse("Transaksi Database Gagal", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Berhasil Membuat Todo", http.StatusOK, "sukses", todo.FormatTodoDetail(newTodo))
	c.JSON(http.StatusOK, response)

}

func (h *todoHandler) UpdateTodo(c *gin.Context) {
	var inputID todo.GetTodoDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Gagal Mengkaitkan ke URI", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData todo.CreateTodoInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Gagal Mengkaitkan Ke JSON", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("userAktif").(user.User)
	inputData.User = currentUser

	updateTodo, err := h.todoService.UpdateTodo(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Gagal Memperbaharui Todo", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Berhasil Memperbaharui Todo", http.StatusOK, "sukses", todo.FormatTodo(updateTodo))
	c.JSON(http.StatusOK, response)
}

func (h *todoHandler) DeleteTodo(c *gin.Context) {
	var inputID todo.GetTodoDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Gagak Mengkaitkan Ke URI", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData todo.CreateTodoInput

	currentUser := c.MustGet("userAktif").(user.User)
	inputData.User = currentUser

	deleteTodo, errr := h.todoService.DeleteTodo(inputID, inputData)

	if errr != nil {
		response := helper.APIResponse("Transaksi Database Gagal", http.StatusUnprocessableEntity, "error", errr)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Berhasil Menghapus Todo", http.StatusOK, "sukses", todo.FormatTodo(deleteTodo))
	c.JSON(http.StatusOK, response)
}
