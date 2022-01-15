package service

import (
	"errors"

	"github.com/heartBrokenGod/direference/models"
	"github.com/heartBrokenGod/direference/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (userService *UserServiceImpl) GetUsers() ([]models.User, error) {
	// no business logic for demo purpose only
	return userService.userRepository.GetUsers()
}

func NewUserServiceImpl(userRepository repository.UserRepository) (*UserServiceImpl, error) {
	// to avoid null ptr reference
	if userRepository == nil {
		return nil, errors.New("userRepository dependency is nil")
	}
	return &UserServiceImpl{
		userRepository: userRepository,
	}, nil
}
