package domain

import (
	"context"
	"time"
)

type User struct {
	UserID          int64     `json:"user_id,omitempty"`
	Name            string    `json:"name,omitempty"`
	Number          string    `json:"number,omitempty"`
	Password        string    `json:"password,omitempty"`
	ConfirmPassword string    `json:"confirm_password,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) (int64, error)
	// TODO update
	GetByID(ctx context.Context, id int64) (*User, error)
	Delete(ctx context.Context, id int64) error
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (int64, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Delete(ctx context.Context, id int64) error
}
