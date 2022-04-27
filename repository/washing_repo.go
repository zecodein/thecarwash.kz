package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/zecodein/thecarwash.kz/domain"
)

type washingRepostiroty struct {
	db *pgxpool.Pool
}

func NewWashingRepostiroty(db *pgxpool.Pool) domain.WashingRepository {
	return &washingRepostiroty{
		db: db,
	}
}

func (w *washingRepostiroty) Create(ctx context.Context, washing *domain.Washing) (int64, error) {
	stmt := `INSERT INTO "washing" ("name", "address", "iin_or_bin", "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5) RETURNING "washing_id"`

	var id int64

	err := w.db.QueryRow(ctx, stmt, washing.Name, washing.Address, washing.IinOrBin, washing.CreatedAt, washing.UpdatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (w *washingRepostiroty) GetByID(ctx context.Context, id int64) (*domain.Washing, error) {
	stmt := `SELECT * FROM "washing" WHERE "washing_id" = $1`
	washing := domain.Washing{}

	row := w.db.QueryRow(ctx, stmt, id)
	err := row.Scan(&washing.WashingID, &washing.Name, &washing.Address, &washing.IinOrBin, &washing.CreatedAt, &washing.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &washing, nil
}

func (w *washingRepostiroty) GetAll(ctx context.Context) (*[]domain.Washing, error) {
	stmt := `SELECT * FROM "washing"`
	washing := domain.Washing{}
	carWashes := []domain.Washing{}

	rows, err := w.db.Query(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&washing.WashingID, &washing.Name, &washing.Address, &washing.IinOrBin, &washing.CreatedAt, &washing.UpdatedAt)
		if err != nil {
			return nil, err
		}

		carWashes = append(carWashes, washing)
	}

	return &carWashes, nil
}

func (w *washingRepostiroty) Delete(ctx context.Context, id int64) error {
	stmt := `DELETE FROM "washing" WHERE "washing_id" = $1`

	_, err := w.db.Exec(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
