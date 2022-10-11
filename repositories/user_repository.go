package repositories

import (
	"context"
	"database/sql"
	"project-fa/models"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, newUser models.CreateUserRequest) error
}

type UserRepository struct {
	mysql *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		mysql: db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, newUser models.CreateUserRequest) error {
	query := "INSERT INTO users(username, email, password, phone_number, age) VALUES (?, ?, ?, ?, ?)"

	_, err := ur.mysql.ExecContext(ctx, query, newUser.Username, newUser.Email, newUser.Password, newUser.PhoneNumber, newUser.Age)
	if err != nil {
		return err
	}

	return nil
}
