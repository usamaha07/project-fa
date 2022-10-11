package models

import "time"

type User struct {
	Id          int       `json:"id" db:"id"`
	Username    string    `json:"username" db:"username"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Age         int       `json:"age" db:"age"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateUserRequest struct {
	Username    string `json:"username" db:"username"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Age         int    `json:"age" db:"age"`
}
