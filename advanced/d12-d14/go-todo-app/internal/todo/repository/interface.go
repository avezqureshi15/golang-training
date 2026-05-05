package todo

import (
	"go-todo-app/internal/platform/db"
	models "go-todo-app/internal/todo/models"
)
type TodoRepository interface {
	db.Repository[models.Todo,int]
}
