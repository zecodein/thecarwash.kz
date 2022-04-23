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
	stmt := `INSERT INTO "user" ("name", "number", "password", "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5) RETURNING "user_id"`

	var id int64

	err := u.db.QueryRow(ctx, stmt, user.Name, user.Number, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return nil, nil
}

func (u *userRepository) GetByNumber(ctx context.Context, phone string) (*domain.User, error) {
	stmt := `SELECT * FROM "user" WHERE "number" = $1`
	user := domain.User{}
	row := u.db.QueryRow(ctx, stmt, phone)

	err := row.Scan(&user.UserID, &user.Name, &user.Number, &user.Access, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) Delete(ctx context.Context, id int64) error {
	return nil
}
