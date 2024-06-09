package helper

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

<<<<<<< HEAD
func APIResponse(message string, code int, status string, data interface{}) Response {
=======
func APIResponse(message string, code int, status string, data interface{}) Response  {

>>>>>>> 9ede165dd324e1863802b8cdb43c54dc29b7457e
	var meta = Meta {
		Message: message,
		Code:    code,
		Status:  status,
	}

	var response = Response {
		Meta: meta,
		Data: data,
	}

	return response
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
