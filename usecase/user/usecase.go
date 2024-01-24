package user

import (
	"errors"
	"go-api-mini-shop/domain"
	"go-api-mini-shop/utils"
)

type UseCase struct {
	Repository domain.UserRepository
}

// Login implements domain.UserUsecase.
func (u *UseCase) Login(user domain.User) (string, error) {
	result, err := u.Repository.FindUserByEmail(user.Email)

	if err != nil {
		return "", err
	}

	// Check password
	if !utils.CheckPasswordHash(user.Password, result.Password) {
		return "", errors.New("Wrong password")
	}

	// Generate token
	dataToken := map[string]interface{}{
		"id":    result.ID,
		"email": result.Email,
		"name":  result.Name,
	}

	token, err := utils.GenerateToken(dataToken)

	if err != nil {
		return "", err
	}

	return token, nil
}

// Register implements domain.UserUsecase.
func (u *UseCase) Register(user domain.User) (*domain.User, error) {
	// Check if email already exists
	_, err := u.Repository.FindUserByEmail(user.Email)

	if err == nil {
		return nil, errors.New("Email already exists")
	}

	// Hash password
	user.Password, err = utils.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	// Create user
	result, err := u.Repository.Create(&user)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewUseCase(repository domain.UserRepository) domain.UserUsecase {
	return &UseCase{
		Repository: repository,
	}
}
