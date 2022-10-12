package repositories

import (
	"context"
	"database/sql"
	"errors"
	"project-fa/models"
)

type AuthRepositoryInterface interface {
	Login(ctx context.Context, email string, password string) (models.UserResponse, error)
}

type AuthRepository struct {
	mysql *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		mysql: db,
	}
}

func (ar *AuthRepository) Login(ctx context.Context, email string, password string) (models.UserResponse, error) {
	query := "SELECT id, username, email, phone_number, age, created_at, updated_at FROM users WHERE email = ? AND password = ?"

	var user models.UserResponse
	err := ar.mysql.QueryRowContext(ctx, query, email, password).Scan(&user.Id, &user.Username, &user.Email, &user.PhoneNumber, &user.Age, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserResponse{}, errors.New("data not found")
		}
		return models.UserResponse{}, err
	}

	return user, err
}
