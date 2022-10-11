package services

import (
	"context"
	"project-fa/models"
	"project-fa/repositories"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, newUser models.CreateUserRequest) error
}

type UserService struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserService(userRepo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepository: userRepo,
	}
}

func (us *UserService) CreateUser(ctx context.Context, newUser models.CreateUserRequest) error {
	err := us.userRepository.CreateUser(ctx, newUser)
	return err
}
