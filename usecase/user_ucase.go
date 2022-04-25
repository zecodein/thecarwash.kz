package usecase

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/zecodein/thecarwash.kz/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

var (
	name, _   = regexp.Compile(`^([А-Я]{1}[а-яё]{1,23}|[A-Z]{1}[a-z]{1,23})$`)
	number, _ = regexp.Compile(`^((\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{7,7}$`)
	err       error
)

func (u *userUsecase) Create(ctx context.Context, user *domain.User) (int64, error) {
	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Number) == "" {
		return 0, domain.ErrInvalidData
	}

	if strings.TrimSpace(user.Password) == "" || strings.TrimSpace(user.ConfirmPassword) == "" {
		return 0, domain.ErrInvalidData
	}

	if user.Password != user.ConfirmPassword {
		return 0, domain.ErrInvalidData
	}

	if !name.MatchString(user.Name) || !number.MatchString(user.Number) {
		return 0, domain.ErrInvalidData
	}

	if len(user.Number) != 12 {
		return 0, domain.ErrInvalidData
	}

	user.Password, err = u.hashPassword(ctx, user.Password)
	if err != nil {
		return 0, err
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	id, err := u.userRepo.Create(ctx, user)
	if err != nil {
		if strings.Contains(err.Error(), "unique") {
			return 0, domain.ErrUniqueData
		}
		return 0, err
	}

	return id, nil
}

func (u *userUsecase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return nil, nil
}

func (u *userUsecase) GetByNumber(ctx context.Context, phone string) (*domain.User, error) {
	if !number.MatchString(phone) || len(phone) != 12 {
		return nil, domain.ErrInvalidData
	}

	user, err := u.userRepo.GetByNumber(ctx, phone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) GetAccess(ctx context.Context, id int64) (string, error) {
	access, err := u.userRepo.GetAccess(ctx, id)
	if err != nil {
		return "", err
	}

	return access, nil
}

func (u *userUsecase) Delete(ctx context.Context, id int64) error {
	return nil
}

func (u *userUsecase) hashPassword(ctx context.Context, password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

func (u *userUsecase) CheckPassword(ctx context.Context, hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
