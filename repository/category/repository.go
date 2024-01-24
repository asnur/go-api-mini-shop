package category

import (
	"go-api-mini-shop/domain"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// FindByID implements domain.CategoryRepository.
func (r *Repository) FindByID(id int) (*domain.Category, error) {
	var category domain.Category

	if err := r.DB.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// Create implements domain.CategoryRepository.
func (r *Repository) Create(category *domain.Category) (*domain.Category, error) {
	if err := r.DB.Create(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

// Delete implements domain.CategoryRepository.
func (r *Repository) Delete(id int) error {
	if err := r.DB.Where("id = ?", id).Delete(&domain.Category{}).Error; err != nil {
		return err
	}

	return nil
}

// FindAll implements domain.CategoryRepository.
func (r *Repository) FindAll() ([]*domain.Category, error) {
	var categories []*domain.Category

	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

// Update implements domain.CategoryRepository.
func (r *Repository) Update(category *domain.Category) (*domain.Category, error) {
	if err := r.DB.Save(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func NewRepository(db *gorm.DB) domain.CategoryRepository {
	return &Repository{
		DB: db,
	}
}
