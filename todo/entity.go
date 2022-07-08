package todo

import (
	"time"
	"todolist/user"
)

type Todo struct {
	ID        int
	Todo      string
	DateTime  time.Time
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time

	User user.User
}
