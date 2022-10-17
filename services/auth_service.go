package services

import (
	"context"
	"project-fa/middlewares"
	"project-fa/models"
	"project-fa/repositories"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, email string, password string) (models.LoginUserResponse, error)
}

type AuthService struct {
	authRepository repositories.AuthRepositoryInterface
}

func NewAuthService(authRepo repositories.AuthRepositoryInterface) AuthServiceInterface {
	return &AuthService{
		authRepository: authRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, email string, password string) (models.LoginUserResponse, error) {
	user, err := as.authRepository.Login(ctx, email, password)
	if err != nil {
		return models.LoginUserResponse{}, err
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
