package order

import (
	"go-api-mini-shop/domain"
)

type UseCase struct {
	Repository     domain.OrderRepository
	RepositoryCart domain.CartRepository
}

// Delete implements domain.OrderUsecase.
func (u *UseCase) Delete(id int) error {
	if err := u.Repository.Delete(id); err != nil {
		return err
	}

	return nil
}

// GetAll implements domain.OrderUsecase.
func (u *UseCase) GetAll(user_id int) ([]*domain.Order, error) {
	result, err := u.Repository.FindAll(user_id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetByID implements domain.OrderUsecase.
func (u *UseCase) GetByID(id int) (*domain.Order, error) {
	result, err := u.Repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Insert implements domain.OrderUsecase.
func (u *UseCase) Insert(user_id int, cart_id []int) (*domain.Order, error) {
	var order domain.Order

	order.UserID = user_id

	// Get total price
	for _, id := range cart_id {
		cart, err := u.RepositoryCart.FindByID(id)

		if err != nil {
			return nil, err
		}

		order.Total += cart.Total

		order.OrderDetails = append(order.OrderDetails, domain.OrderDetail{
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Total:     cart.Total,
		})
	}

	result, err := u.Repository.Create(&order)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewUseCase(repository domain.OrderRepository, repositoryCart domain.CartRepository) domain.OrderUsecase {
	return &UseCase{
		Repository:     repository,
		RepositoryCart: repositoryCart,
	}
}
