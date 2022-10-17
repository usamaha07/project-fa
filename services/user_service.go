package services

import (
	"context"
	"errors"
	"project-fa/models"
	"project-fa/repositories"
	"time"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, newUser models.CreateUserRequest) error
	GetUserById(ctx context.Context, userId int) (models.UserResponse, error)
	GetAllUser(ctx context.Context) ([]models.UserResponse, error)
	DeleteUser(ctx context.Context, idToken int) error
	UpdateUser(ctx context.Context, updateUser models.UpdateUserRequest, idToken int) (models.UserResponse, error)
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

	if newUser.Username == "" {
		return errors.New("username is required")
	}

	if newUser.Email == "" {
		return errors.New("email is required")
	}

	if newUser.PhoneNumber == "" {
		return errors.New("phone number is required")
	}

	if newUser.Age == 0 {
		return errors.New("age is required")
	}

	if newUser.Password == "" {
		return errors.New("password is required")
	}

	err := us.userRepository.CreateUser(ctx, newUser)
	return err
}

func (us *UserService) GetUserById(ctx context.Context, userId int) (models.UserResponse, error) {
	user, err := us.userRepository.GetUserById(ctx, userId)

	userResponse := models.UserResponse{
		Id:          user.Id,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Age:         user.Age,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return userResponse, err
}

func (us *UserService) GetAllUser(ctx context.Context) ([]models.UserResponse, error) {
	users, err := us.userRepository.GetAllUser(ctx)
	return users, err
}

func (us *UserService) DeleteUser(ctx context.Context, idToken int) error {
	err := us.userRepository.DeleteUser(ctx, idToken)
	return err
}

func (us *UserService) UpdateUser(ctx context.Context, updateUser models.UpdateUserRequest, idToken int) (models.UserResponse, error) {
	getUser, err := us.userRepository.GetUserById(ctx, idToken)
	if err != nil {
		return models.UserResponse{}, err
	}

	if updateUser.Username != "" {
		getUser.Username = updateUser.Username
	}

	if updateUser.Email != "" {
		getUser.Email = updateUser.Email
	}

	if updateUser.Password != "" {
		getUser.Password = updateUser.Password
	}

	if updateUser.PhoneNumber != "" {
		getUser.PhoneNumber = updateUser.PhoneNumber
	}

	if updateUser.Age != 0 {
		getUser.Age = updateUser.Age
	}

	layoutFormat := "2006-01-02T15:04:05"
	value := time.Now().Local().Format("2006-01-02T15:04:05")

	now, _ := time.Parse(layoutFormat, value)
	getUser.UpdatedAt = &now

	user, err := us.userRepository.UpdateUser(ctx, getUser, idToken)

	responseUpdate := models.UserResponse{
		Id:          getUser.Id,
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Age:         user.Age,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return responseUpdate, err
}
