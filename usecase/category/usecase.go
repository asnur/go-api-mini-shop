package category

import (
	"go-api-mini-shop/domain"
	"go-api-mini-shop/utils"
)

type UseCase struct {
	Repository domain.CategoryRepository
}

// Delete implements domain.CategoryUsecase.
func (u *UseCase) Delete(id int) error {
	if err := u.Repository.Delete(id); err != nil {
		return err
	}

	return nil
}

// Detail implements domain.CategoryUsecase.
func (u *UseCase) Detail(id int) (*domain.Category, error) {
	result, err := u.Repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAll implements domain.CategoryUsecase.
func (u *UseCase) GetAll() ([]*domain.Category, error) {
	result, err := u.Repository.FindAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Insert implements domain.CategoryUsecase.
func (u *UseCase) Insert(name string) (*domain.Category, error) {
	var category domain.Category

	category.Name = name
	category.Slug = utils.SlugString(name)

	result, err := u.Repository.Create(&category)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Update implements domain.CategoryUsecase.
func (u *UseCase) Update(id int, name string) (*domain.Category, error) {
	var category domain.Category

	category.ID = id
	category.Name = name
	category.Slug = utils.SlugString(name)

	result, err := u.Repository.Update(&category)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewUseCase(respository domain.CategoryRepository) domain.CategoryUsecase {
	return &UseCase{
		Repository: respository,
	}
}
