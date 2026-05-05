package todo

import (
	models "go-todo-app/internal/todo/models"
	appErr "go-todo-app/pkg/errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TodoRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewTodoRepo(db *gorm.DB, log *zap.Logger) *TodoRepo {
	return &TodoRepo{
		db:  db,
		log: log,
	}
}

// GetAll retrieves all todos
func (r *TodoRepo) GetAll() ([]models.Todo, error) {
	var todos []models.Todo

	if err := r.db.Find(&todos).Error; err != nil {
		r.log.Error("failed to fetch todos", zap.Error(err))
		return nil, err
	}

	r.log.Info("fetched todos", zap.Int("count", len(todos)))
	return todos, nil
}

// GetByID retrieves a todo by ID
func (r *TodoRepo) GetByID(id int) (models.Todo, error) {
	var todo models.Todo

	err := r.db.First(&todo, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.log.Warn("todo not found", zap.Int("id", id))
			return models.Todo{}, appErr.ErrNotFound
		}

		r.log.Error("failed to fetch todo", zap.Int("id", id), zap.Error(err))
		return models.Todo{}, err
	}

	r.log.Info("fetched todo", zap.Int("id", id))
	return todo, nil
}

// Create inserts a new todo
func (r *TodoRepo) Create(todo models.Todo) (models.Todo, error) {
	if err := r.db.Create(&todo).Error; err != nil {
		r.log.Error("failed to create todo",
			zap.String("title", todo.Title),
			zap.Error(err),
		)
		return models.Todo{}, err
	}

	r.log.Info("created todo",
		zap.Int("id", todo.ID),
		zap.String("title", todo.Title),
	)

	return todo, nil
}

// Update modifies an existing todo
func (r *TodoRepo) Update(id int, updated models.Todo) (models.Todo, error) {
	var todo models.Todo

	// check existence
	if err := r.db.First(&todo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.log.Warn("todo not found for update", zap.Int("id", id))
			return models.Todo{}, appErr.ErrNotFound
		}
		return models.Todo{}, err
	}

	// update fields
	todo.Title = updated.Title
	todo.Done = updated.Done

	if err := r.db.Save(&todo).Error; err != nil {
		r.log.Error("failed to update todo",
			zap.Int("id", id),
			zap.Error(err),
		)
		return models.Todo{}, err
	}

	r.log.Info("updated todo",
		zap.Int("id", id),
		zap.String("title", todo.Title),
		zap.Bool("done", todo.Done),
	)

	return todo, nil
}

// Delete removes a todo
func (r *TodoRepo) Delete(id int) error {
	result := r.db.Delete(&models.Todo{}, id)

	if result.Error != nil {
		r.log.Error("failed to delete todo",
			zap.Int("id", id),
			zap.Error(result.Error),
		)
		return result.Error
	}

	if result.RowsAffected == 0 {
		r.log.Warn("todo not found for delete", zap.Int("id", id))
		return appErr.ErrNotFound
	}

	r.log.Info("deleted todo", zap.Int("id", id))
	return nil
}