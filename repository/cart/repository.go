package cart

import (
	"go-api-mini-shop/domain"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// FindByID implements domain.CartRepository.
func (r *Repository) FindByID(id int) (*domain.Cart, error) {
	var cart domain.Cart

	if err := r.DB.Preload("Product.Category").Preload("User").First(&cart, id).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

// Create implements domain.CartRepository.
func (r *Repository) Create(cart *domain.Cart) (*domain.Cart, error) {
	if err := r.DB.Create(&cart).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Preload("Product.Category").Preload("User").First(&cart, cart.ID).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

// Delete implements domain.CartRepository.
func (r *Repository) Delete(id int) error {
	if err := r.DB.Where("id = ?", id).Delete(&domain.Cart{}).Error; err != nil {
		return err
	}

	return nil
}

// FindAll implements domain.CartRepository.
func (r *Repository) FindAll(params map[string]any) ([]*domain.Cart, error) {
	var carts []*domain.Cart

	if err := r.DB.Preload("Product.Category").Preload("User").Where(params).Find(&carts).Error; err != nil {
		return nil, err
	}

	return carts, nil
}

// Update implements domain.CartRepository.
func (r *Repository) Update(cart *domain.Cart) (*domain.Cart, error) {
	if err := r.DB.Save(&cart).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Preload("Product.Category").Preload("User").First(&cart, cart.ID).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

func NewRepository(db *gorm.DB) domain.CartRepository {
	return &Repository{
		DB: db,
	}
}
