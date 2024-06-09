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
	return &todoHandler {
		todoService,
		authService,
	}
}


func (h *todoHandler) GetTodos(c *gin.Context) {
<<<<<<< HEAD
	// TODO
	// 1. konversi nilai dari params user_id
	// 2. ambil data list todo dari method GetTodos
	// 3. Return data todo ke bentuk JSON
	
	// TODO 1
	userID, _ := strconv.Atoi(c.Query("user_id"))
	
	// TODO 2
	todos, err := h.todoService.GetTodos(userID)
=======

	var currentUser = c.MustGet("who_is_logged_in").(user.User)
	
	var userID = currentUser.ID

	var _page_number_, _ = strconv.Atoi(c.Query("page"))
	
	var todos, err = h.todoService.GetTodos(userID, _page_number_)
>>>>>>> 9ede165dd324e1863802b8cdb43c54dc29b7457e
	if err != nil {
		var _output_ = helper.APIResponse("Gagal Mengambil Data Todos", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, _output_)
		return
	}
<<<<<<< HEAD
	
	// TODO 3
	response := helper.APIResponse("Daftar Todos", http.StatusOK, "sukses", todo.FormatTodos(todos))
	c.JSON(http.StatusOK, response)
=======

	var pagination, _ = h.todoService.GetNumberPaginationOfTotalTodo(userID)
>>>>>>> 9ede165dd324e1863802b8cdb43c54dc29b7457e


	var todos_data = make(map[string]interface{})
	todos_data["number_of_pagination"] = pagination["number_of_pagination"]
	todos_data["total_data"] = pagination["total_data"]
	todos_data["todos"] = todo.FormatTodos(todos)

	var _output_ = helper.APIResponse("Daftar Todos", http.StatusOK, "sukses", todos_data)
	c.JSON(http.StatusOK, _output_)
}


func (h *todoHandler) GetTodo(c *gin.Context) {
<<<<<<< HEAD
	// TODO 
	// 1. akses struct GetTodoDetailInput di package todo
	// 2. mengcocokan json data dengan struct GetTodoDetailInput 
	
	// TODO 1
	var input todo.GetTodoDetailInput

	// TODO 2
	err := c.ShouldBindUri(&input)
=======
	
	var input todo.GetTodoDetailInput
	
	var err = c.ShouldBindUri(&input)
>>>>>>> 9ede165dd324e1863802b8cdb43c54dc29b7457e
	if err != nil {
		var _output_ = helper.APIResponse("Gagal Mengkaitkan ke JSON ", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, _output_)
		return
	}
<<<<<<< HEAD
	
	// TODO 3
	todoDetail, err := h.todoService.GetTodoByID(input)
	if err != nil {
		response := helper.APIResponse("Transaksi Database Gagal", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
=======

	var currentUser = c.MustGet("who_is_logged_in").(user.User)	
	
	var todoDetail, err2 = h.todoService.GetTodoByID(input, currentUser.ID)
	if err2 != nil {
		var _output_ = helper.APIResponse("Transaksi Database Gagal", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, _output_)
>>>>>>> 9ede165dd324e1863802b8cdb43c54dc29b7457e
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

	var currentUser = c.MustGet("who_is_logged_in").(user.User)
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

	var currentUser = c.MustGet("who_is_logged_in").(user.User)
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
		var _output_ = helper.APIResponse("Gagal Mengkaitkan Ke URI", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, _output_)
		return
	}

	var inputData todo.CreateTodoInput

	var currentUser = c.MustGet("who_is_logged_in").(user.User)
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

