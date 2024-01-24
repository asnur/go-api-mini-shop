package cart

import (
	"go-api-mini-shop/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	DB *gorm.DB
}

// Create implements domain.CartRepository.
func (r *Repository) Create(cart *domain.Cart) (*domain.Cart, error) {
	if err := r.DB.Create(&cart).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Preload(clause.Associations).First(&cart, cart.ID).Error; err != nil {
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
func (r *Repository) FindAll(user_id int) ([]*domain.Cart, error) {
	var carts []*domain.Cart

	if err := r.DB.Preload(clause.Associations).Where("user_id = ?", user_id).Find(&carts).Error; err != nil {
		return nil, err
	}

	return carts, nil
}

// Update implements domain.CartRepository.
func (r *Repository) Update(cart *domain.Cart) (*domain.Cart, error) {
	if err := r.DB.Save(&cart).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Preload(clause.Associations).First(&cart, cart.ID).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

func NewRepository(db *gorm.DB) domain.CartRepository {
	return &Repository{
		DB: db,
	}
}
