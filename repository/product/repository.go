package product

import (
	"go-api-mini-shop/domain"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// Create implements domain.ProductRepository.
func (r *Repository) Create(product *domain.Product) (*domain.Product, error) {
	if err := r.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Preload("Category").First(&product, product.ID).Error; err != nil {
		return nil, err
	}

	return product, nil
}

// Delete implements domain.ProductRepository.
func (r *Repository) Delete(id int) error {
	if err := r.DB.Delete(&domain.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}

// FindAll implements domain.ProductRepository.
func (r *Repository) FindAll() ([]*domain.Product, error) {
	var products []*domain.Product

	if err := r.DB.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

// FindByID implements domain.ProductRepository.
func (r *Repository) FindByID(id int) (*domain.Product, error) {
	var product domain.Product

	if err := r.DB.Preload("Category").First(&product, id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

// Update implements domain.ProductRepository.
func (r *Repository) Update(product *domain.Product) (*domain.Product, error) {
	if err := r.DB.Save(&product).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Preload("Category").First(&product, product.ID).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func NewRepository(db *gorm.DB) domain.ProductRepository {
	return &Repository{
		DB: db,
	}
}
