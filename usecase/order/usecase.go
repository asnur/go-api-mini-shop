package order

import "go-api-mini-shop/domain"

type UseCase struct {
	Repository domain.OrderRepository
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
func (u *UseCase) Insert(user_id int, total int, status int) (*domain.Order, error) {
	var order domain.Order

	order.UserID = user_id
	order.Total = total
	order.Status = status

	result, err := u.Repository.Create(&order)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewUseCase(repository domain.OrderRepository) domain.OrderUsecase {
	return &UseCase{
		Repository: repository,
	}
}
