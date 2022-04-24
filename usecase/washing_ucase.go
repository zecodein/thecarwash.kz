package usecase

import (
	"context"

	"github.com/zecodein/thecarwash.kz/domain"
)

type washingUsecase struct {
	washingRepo domain.WashingRepository
}

func NewWashingUsecase(w domain.WashingRepository) domain.WashingUsecase {
	return &washingUsecase{
		washingRepo: w,
	}
}

func (w *washingUsecase) Create(ctx context.Context, washing *domain.Washing) (int64, error) {
	return 0, nil
}

func (w *washingUsecase) GetByID(ctx context.Context, id int64) (*domain.Washing, error) {
	return nil, nil
}

func (w *washingUsecase) GetAll(ctx context.Context) (*[]domain.Washing, error) {
	return nil, nil
}

func (w *washingUsecase) Delete(ctx context.Context, id int64) error {
	return nil
}
