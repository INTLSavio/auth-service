package services

import (
	"github.com/INTLSavio/auth-service/internal/entities"
	"github.com/INTLSavio/auth-service/internal/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (userService *UserService) CreteUser(name, email, phone, address, password string) (*entities.User, error) {
	user, err := entities.NewUser(name, email, phone, address, password)
	if err != nil {
		return nil, err
	}

	err = userService.UserRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
