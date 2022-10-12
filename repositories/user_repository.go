package repositories

import (
	"context"
	"database/sql"
	"errors"
	"project-fa/models"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, newUser models.CreateUserRequest) error
	GetUserById(ctx context.Context, userId int) (models.UserResponse, error)
	GetAllUser(ctx context.Context) ([]models.UserResponse, error)
	DeleteUser(ctx context.Context, idToken int) error
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

func (ur *UserRepository) GetUserById(ctx context.Context, userId int) (models.UserResponse, error) {
	query := "SELECT id, username, email, phone_number, age, created_at, updated_at FROM users WHERE id = ?"

	var user models.UserResponse
	err := ur.mysql.QueryRowContext(ctx, query, userId).Scan(&user.Id, &user.Username, &user.Email, &user.PhoneNumber, &user.Age, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserResponse{}, errors.New("data not found")
		}
		return models.UserResponse{}, err
	}

	return user, nil
}

func (ur *UserRepository) GetAllUser(ctx context.Context) ([]models.UserResponse, error) {
	query := "SELECT id, username, email, phone_number, age, created_at, updated_at FROM users"

	rows, err := ur.mysql.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserResponse
	for rows.Next() {
		var user models.UserResponse
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.PhoneNumber, &user.Age, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, idToken int) error {
	query := "DELETE FROM users WHERE id = ?"

	result, err := ur.mysql.ExecContext(ctx, query, idToken)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("data not found")
	}

	return nil
}
