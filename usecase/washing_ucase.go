package usecase

import (
	"context"
	"regexp"
	"strings"
	"time"

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
	if strings.TrimSpace(washing.Name) == "" || strings.TrimSpace(washing.Address) == "" {
		return 0, domain.ErrInvalidData
	}

	if strings.TrimSpace(washing.Iin) == "" || strings.TrimSpace(washing.Bin) == "" {
		return 0, domain.ErrInvalidData
	}

	rIin, err := regexp.Compile(`^((0[48]|[2468][048]|[13579][26])0229[1-6]|000229[34]|\d\d((0[13578]|1[02])(0[1-9]|[12]\d|3[01])|(0[469]|11)(0[1-9]|[12]\d|30)|02(0[1-9]|1\d|2[0-8]))[1-6])\d{5}$`)
	if err != nil {
		return 0, err
	}

	if !rIin.MatchString(washing.Iin) {
		return 0, domain.ErrInvalidIin
	}

	washing.CreatedAt = time.Now()
	washing.UpdatedAt = time.Now()

	id, err := w.washingRepo.Create(ctx, washing)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (w *washingUsecase) GetByID(ctx context.Context, id int64) (*domain.Washing, error) {
	washing, err := w.washingRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return washing, nil
}

func (w *washingUsecase) GetAll(ctx context.Context) (*[]domain.Washing, error) {
	return nil, nil
}

func (w *washingUsecase) Delete(ctx context.Context, id int64) error {
	return nil
}
