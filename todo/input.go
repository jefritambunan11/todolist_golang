package todo

import (
	"time"
	"todolist/user"
)

type GetTodoDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateTodoInput struct {
	Todo     string    `json:"todo" binding:"required"`
	DateTime time.Time `json:"date_time" binding:"required"`

	User user.User
}
