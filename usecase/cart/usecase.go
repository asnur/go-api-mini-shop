package cart

import "go-api-mini-shop/domain"

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
func (u *UseCase) GetAll(user_id int) ([]*domain.Cart, error) {
	result, err := u.Repository.FindAll(user_id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Insert implements domain.CartUsecase.
func (u *UseCase) Insert(product_id int, user_id int, quantity int) (*domain.Cart, error) {
	var cart domain.Cart

	cart.ProductID = product_id
	cart.UserID = user_id
	cart.Quantity = quantity

	// Get product price
	product, err := u.RepositoryProduct.FindByID(product_id)

	if err != nil {
		return nil, err
	}

	cart.Total = product.Price * quantity

	result, err := u.Repository.Create(&cart)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Update implements domain.CartUsecase.
func (u *UseCase) Update(id int, quantity int) (*domain.Cart, error) {
	var cart domain.Cart

	cart.ID = id
	cart.Quantity = quantity

	// Get product price
	product, err := u.RepositoryProduct.FindByID(cart.ProductID)

	if err != nil {
		return nil, err
	}

	cart.Total = product.Price * quantity

	result, err := u.Repository.Update(&cart)

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
