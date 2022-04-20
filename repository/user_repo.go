package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/zecodein/thecarwash.kz/domain"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) (int64, error) {
	return 0, nil
}

func (u *userRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return nil, nil
}

func (u *userRepository) Delete(ctx context.Context, id int64) error {
	return nil
}
