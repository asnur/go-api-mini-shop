package cart

import (
	"go-api-mini-shop/domain"
)

type UseCase struct {
	Repository        domain.CartRepository
	RepositoryProduct domain.ProductRepository
}

// Delete implements domain.CartUsecase.
func (u *UseCase) Delete(id int) error {
	if err := u.Repository.Delete(id); err != nil {
		return err
	}

	return nil
}

// GetAll implements domain.CartUsecase.
func (u *UseCase) GetAll(params map[string]any) ([]*domain.Cart, error) {
	result, err := u.Repository.FindAll(params)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Insert implements domain.CartUsecase.
func (u *UseCase) Insert(cart *domain.Cart) (*domain.Cart, error) {
	// Check if product already in cart
	productExist, err := u.Repository.FindAll(map[string]any{
		"product_id": cart.ProductID,
		"user_id":    cart.UserID,
	})

	if err != nil {
		return nil, err
	}

	if len(productExist) > 0 {
		product := *productExist[0]

		cart.ID = product.ID
		cart.Quantity = cart.Quantity + product.Quantity
		cart.Total = cart.Quantity * product.Product.Price

		result, err := u.Repository.Update(cart)

		if err != nil {
			return nil, err
		}

		return result, nil
	}

	// Get product price
	product, err := u.RepositoryProduct.FindByID(cart.ProductID)

	if err != nil {
		return nil, err
	}

	cart.Total = product.Price * cart.Quantity

	result, err := u.Repository.Create(cart)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Update implements domain.CartUsecase.
func (u *UseCase) Update(id int, quantity int) (*domain.Cart, error) {
	cart, err := u.Repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	cart.Quantity = quantity
	cart.Total = cart.Quantity * cart.Product.Price

	result, err := u.Repository.Update(cart)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewUseCase(repository domain.CartRepository, repositoryProduct domain.ProductRepository) domain.CartUsecase {
	return &UseCase{
		Repository:        repository,
		RepositoryProduct: repositoryProduct,
	}
}
