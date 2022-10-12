package models

import "time"

type User struct {
	Id          int        `json:"id" db:"id"`
	Username    string     `json:"username" db:"username"`
	Email       string     `json:"email" db:"email"`
	Password    string     `json:"password" db:"password"`
	PhoneNumber string     `json:"phone_number" db:"phone_number"`
	Age         int        `json:"age" db:"age"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}

type UserResponse struct {
	Id          int        `json:"id" db:"id"`
	Username    string     `json:"username" db:"username"`
	Email       string     `json:"email" db:"email"`
	PhoneNumber string     `json:"phone_number" db:"phone_number"`
	Age         int        `json:"age" db:"age"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}

type CreateUserRequest struct {
	Username    string `json:"username" db:"username"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Age         int    `json:"age" db:"age"`
}

type LoginUserRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginUserResponse struct {
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}
