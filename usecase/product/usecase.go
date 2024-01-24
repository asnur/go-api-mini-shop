package product

import (
	"go-api-mini-shop/domain"
)

type UseCase struct {
	Repository domain.ProductRepository
}

// Create implements domain.ProductUsecase.
func (u *UseCase) Create(product domain.Product) (*domain.Product, error) {
	result, err := u.Repository.Create(&product)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete implements domain.ProductUsecase.
func (u *UseCase) Delete(id int) error {
	if err := u.Repository.Delete(id); err != nil {
		return err
	}

	return nil
}

// GetAll implements domain.ProductUsecase.
func (u *UseCase) GetAll(params map[string]any) ([]*domain.Product, error) {

	result, err := u.Repository.FindAll(params)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetByID implements domain.ProductUsecase.
func (u *UseCase) GetByID(id int) (*domain.Product, error) {
	result, err := u.Repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Update implements domain.ProductUsecase.
func (u *UseCase) Update(product domain.Product) (*domain.Product, error) {
	result, err := u.Repository.Update(&product)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewUseCase(repository domain.ProductRepository) domain.ProductUsecase {
	return &UseCase{
		Repository: repository,
	}
}
