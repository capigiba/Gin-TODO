package repository

import (
	"note/internal/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAll(userID uint64) ([]models.Todo, error)
	Create(todo *models.Todo) error
	Find(userID, id uint64) (*models.Todo, error)
	Update(todo *models.Todo) error
}

type GormTodoRepository struct {
	db *gorm.DB
}

func NewGormTodoRepository(db *gorm.DB) *GormTodoRepository {
	return &GormTodoRepository{
		db: db,
	}
}

func (repo *GormTodoRepository) GetAll(userID uint64) ([]models.Todo, error) {
	var todos []models.Todo
	err := repo.db.Where(`deleted_at IS NULL AND user_id = ?`, userID).Order(`updated_at desc`).Find(&todos).Error
	return todos, err
}

func (repo *GormTodoRepository) Create(todo *models.Todo) error {
	return repo.db.Create(todo).Error
}

func (repo *GormTodoRepository) Find(userID, id uint64) (*models.Todo, error) {
	var todo models.Todo
	err := repo.db.Where(`id = ? AND user_id = ?`, id, userID).Find(&todo).Error
	return &todo, err
}

func (repo *GormTodoRepository) Update(todo *models.Todo) error {
	return repo.db.Save(todo).Error
}
