package repositories

import (
	"context"
	"database/sql"
	"errors"
	"project-fa/models"
)

type AuthRepositoryInterface interface {
	Login(ctx context.Context, email string) (models.LoginDataUserResponse, error)
}

type AuthRepository struct {
	mysql *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		mysql: db,
	}
}

func (ar *AuthRepository) Login(ctx context.Context, email string) (models.LoginDataUserResponse, error) {
	query := "SELECT id, email, password FROM users WHERE email = ?"

	var user models.LoginDataUserResponse
	err := ar.mysql.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.LoginDataUserResponse{}, errors.New("data not found")
		}
		return models.LoginDataUserResponse{}, err
	}

	return user, err
}
