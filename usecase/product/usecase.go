package product

import (
	"go-api-mini-shop/domain"
	repositoryRedis "go-api-mini-shop/repository/product"
)

type UseCase struct {
	Repository      domain.ProductRepository
	RepositoryRedis repositoryRedis.RepositoryRedis
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
	// Check if category ID provided
	categoryID, ok := params["category_id"].(float64)

	if ok {
		products, err := u.RepositoryRedis.GetProductByCategoryID(int(categoryID))

		if err != nil || len(products) == 0 {
			// If data not found in redis or there's an error, fetch from the main repository
			result, err := u.Repository.FindAll(params)

			if err != nil {
				return nil, err
			}

			// Set data to redis
			if err := u.RepositoryRedis.SetProductByCategoryID(int(categoryID), result); err != nil {
				return nil, err
			}

			return result, nil
		}

		return products, nil
	}

	// No category ID provided, fetch from the main repository
	products, err := u.Repository.FindAll(params)

	if err != nil {
		return nil, err
	}

	return products, nil
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

func NewUseCase(repository domain.ProductRepository, repositoryRedis repositoryRedis.RepositoryRedis) domain.ProductUsecase {
	return &UseCase{
		Repository:      repository,
		RepositoryRedis: repositoryRedis,
	}
}
