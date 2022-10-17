package services

import (
	"context"
	"errors"
	"project-fa/helpers"
	"project-fa/middlewares"
	"project-fa/models"
	"project-fa/repositories"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, userLogin models.LoginUserRequest) (models.LoginUserResponse, error)
}

type AuthService struct {
	authRepository repositories.AuthRepositoryInterface
}

func NewAuthService(authRepo repositories.AuthRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		authRepository: authRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, userLogin models.LoginUserRequest) (models.LoginUserResponse, error) {

	if userLogin.Email == "" {
		return models.LoginUserResponse{}, errors.New("email is required")
	}

	user, err := as.authRepository.Login(ctx, userLogin.Email)
	if err != nil {
		return models.LoginUserResponse{}, err
	}

	if !helpers.CheckPassHash(userLogin.Password, user.Password) {
		return models.LoginUserResponse{}, errors.New("password incorrect")
	}

	token, err := middlewares.CreateToken(user.Id)
	if err != nil {
		return models.LoginUserResponse{}, err
	}

	loginResponse := models.LoginUserResponse{
		Token:  token,
		UserId: user.Id,
	}

	return loginResponse, err
}
