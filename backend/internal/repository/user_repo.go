package repository

import (
	"note/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(email string) (*models.User, error)
	GetByID(id uint64) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(user *models.User) error
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (repo *GormUserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (repo *GormUserRepository) GetByID(id uint64) (*models.User, error) {
	var user models.User
	err := repo.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (repo *GormUserRepository) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *GormUserRepository) Update(user *models.User) error {
	return repo.db.Save(user).Error
}

func (repo *GormUserRepository) Delete(user *models.User) error {
	return repo.db.Delete(user).Error
}
