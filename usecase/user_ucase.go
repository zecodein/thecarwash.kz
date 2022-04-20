package usecase

import (
	"context"

	"github.com/zecodein/thecarwash.kz/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) Create(ctx context.Context, user *domain.User) (int64, error) {
	return 0, nil
}

func (u *userUsecase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return nil, nil
}

func (u *userUsecase) Delete(ctx context.Context, id int64) error {
	return nil
}
