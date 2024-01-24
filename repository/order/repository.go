package order

import (
	"go-api-mini-shop/domain"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// Create implements domain.OrderRepository.
func (r *Repository) Create(order *domain.Order) (*domain.Order, error) {
	if err := r.DB.Create(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

// Delete implements domain.OrderRepository.
func (r *Repository) Delete(id int) error {
	if err := r.DB.Delete(&domain.Order{}, id).Error; err != nil {
		return err
	}

	return nil
}

// FindAll implements domain.OrderRepository.
func (r *Repository) FindAll(user_id int) ([]*domain.Order, error) {
	var orders []*domain.Order

	if err := r.DB.Preload("OrderItems").Where("user_id", user_id).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

// FindByID implements domain.OrderRepository.
func (r *Repository) FindByID(id int) (*domain.Order, error) {
	var order domain.Order

	if err := r.DB.Preload("OrderItems").First(&order, id).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func NewRepository(db *gorm.DB) domain.OrderRepository {
	return &Repository{
		DB: db,
	}
}
