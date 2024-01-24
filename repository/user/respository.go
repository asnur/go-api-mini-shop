package user

import (
	"go-api-mini-shop/domain"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// Delete implements domain.UserRepository.
func (r *Repository) Delete(id int) error {
	if err := r.DB.Delete(&domain.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

// FindByID implements domain.UserRepository.
func (r *Repository) FindByID(id int) (*domain.User, error) {
	var user domain.User

	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Update implements domain.UserRepository.
func (r *Repository) Update(user *domain.User) (*domain.User, error) {
	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Create implements domain.UserRepository.
func (r *Repository) Create(user *domain.User) (*domain.User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// FindUserByEmail implements domain.UserRepository.
func (r *Repository) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User

	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func NewRepository(db *gorm.DB) domain.UserRepository {
	return &Repository{
		DB: db,
	}
}
