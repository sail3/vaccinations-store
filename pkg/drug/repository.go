package drug

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Repository interface {
	RegisterDrug(context.Context, Drug) (int, error)
	ListDrug(context.Context) ([]Drug, error)
	UpdateDrug(context.Context, int, Drug) (Drug, error)
	DeleteDrug(context.Context, int) error
}

func NewRepository(p *sql.DB) Repository {
	return &repository{
		db: p,
	}
}

type repository struct {
	db *sql.DB
}

const maxTimeout = 15

func (r repository) RegisterDrug(ctx context.Context, d Drug) (int, error) {
	q := `INSERT INTO drug (name, approved, min_dose, max_dose, available_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	ctx, cancel := context.WithTimeout(ctx, maxTimeout*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, q)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var lastID int
	err = stmt.QueryRowContext(ctx,
		d.Name,
		d.Approved,
		d.MinDose,
		d.MaxDose,
		d.AvailableAt,
	).Scan(&lastID)
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func (r repository) ListDrug(ctx context.Context) ([]Drug, error) {
	q := `SELECT id, name, approved, min_dose, max_dose, available_at FROM drug ORDER BY id ASC`
	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}
	res := make([]Drug, 0)
	for rows.Next() {
		var d Drug
		rows.Scan(
			&d.ID,
			&d.Name,
			&d.Approved,
			&d.MinDose,
			&d.MaxDose,
			&d.AvailableAt,
		)
		res = append(res, d)
	}
	return res, nil
}

func (r repository) UpdateDrug(ctx context.Context, id int, d Drug) (Drug, error) {
	q := `UPDATE drug SET name=$1, approved=$2, min_dose=$3, max_dose=$4, available_at=$5 WHERE id=$6`
	stmt, err := r.db.PrepareContext(ctx, q)
	if err != nil {
		return Drug{}, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx,
		d.Name,
		d.Approved,
		d.MinDose,
		d.MaxDose,
		d.AvailableAt,
		id,
	)
	if err != nil {
		return Drug{}, err
	}
	ra, err := result.RowsAffected()
	if err != nil {
		return Drug{}, err
	} else if ra != 1 {
		return Drug{}, errors.New("expected 1 row affected")
	}
	d.ID = id
	return d, nil
}

func (r repository) DeleteDrug(ctx context.Context, id int) error {
	q := `DELETE FROM drug WHERE id = $1`
	stmt, err := r.db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	if val, err := res.RowsAffected(); val < 1 || err != nil {
		return errors.New("row with provided id doesn't exists")
	}

	return nil
}
