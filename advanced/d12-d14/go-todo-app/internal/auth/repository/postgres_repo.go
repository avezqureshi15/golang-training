package auth

import (
	models "go-todo-app/internal/auth/models"

	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepository {
	return &authRepo{db: db}
}

func (r *authRepo) Create(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *authRepo) GetByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *authRepo) GetAll() ([]models.User, error) {
	var users []models.User

	err := r.db.Select("id", "name", "email").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}