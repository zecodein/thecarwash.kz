package domain

import (
	"context"
	"time"
)

type Washing struct {
	WashingID int64     `json:"washing_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Address   string    `json:"address,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type WashingUsecase interface {
	Create(ctx context.Context, washing *Washing) (int64, error)
	GetByID(ctx context.Context, id int64) (*Washing, error)
	GetAll(ctx context.Context) (*[]Washing, error)
	Delete(ctx context.Context, id int64) error
}

type WashingRepository interface {
	Create(ctx context.Context, washing *Washing) (int64, error)
	GetByID(ctx context.Context, id int64) (*Washing, error)
	GetAll(ctx context.Context) (*[]Washing, error)
	Delete(ctx context.Context, id int64) error
}
