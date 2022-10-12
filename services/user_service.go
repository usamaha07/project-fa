package services

import (
	"context"
	"project-fa/models"
	"project-fa/repositories"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, newUser models.CreateUserRequest) error
	GetUserById(ctx context.Context, userId int) (models.UserResponse, error)
	GetAllUser(ctx context.Context) ([]models.UserResponse, error)
	DeleteUser(ctx context.Context, idToken int) error
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

func (us *UserService) GetUserById(ctx context.Context, userId int) (models.UserResponse, error) {
	user, err := us.userRepository.GetUserById(ctx, userId)
	return user, err
}

func (us *UserService) GetAllUser(ctx context.Context) ([]models.UserResponse, error) {
	users, err := us.userRepository.GetAllUser(ctx)
	return users, err
}

func (us *UserService) DeleteUser(ctx context.Context, idToken int) error {
	err := us.userRepository.DeleteUser(ctx, idToken)
	return err
}
