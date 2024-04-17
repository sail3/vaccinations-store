package vaccination

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Repository interface {
	RegisterVaccination(context.Context, Vaccination) (int, error)
	ListVaccination(context.Context) ([]Vaccination, error)
	UpdateVaccination(context.Context, int, Vaccination) (Vaccination, error)
	DeleteVaccination(context.Context, int) error
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

func (r repository) RegisterVaccination(ctx context.Context, d Vaccination) (int, error) {
	q := `INSERT INTO vaccination (name, drug_id, dose, date) VALUES ($1, $2, $3, $4) RETURNING id `
	ctx, cancel := context.WithTimeout(ctx, maxTimeout*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, q)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	var id int
	err = stmt.QueryRowContext(ctx,
		d.Name,
		d.DrugID,
		d.Dose,
		d.Date,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r repository) ListVaccination(ctx context.Context) ([]Vaccination, error) {
	q := `SELECT id, name, drug_id, dose, date FROM vaccination ORDER BY id ASC`
	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}
	res := make([]Vaccination, 0)
	for rows.Next() {
		var d Vaccination
		rows.Scan(
			&d.ID,
			&d.Name,
			&d.DrugID,
			&d.Dose,
			&d.Date,
		)
		res = append(res, d)
	}
	return res, nil
}

func (r repository) UpdateVaccination(ctx context.Context, id int, d Vaccination) (Vaccination, error) {
	q := `UPDATE vaccination SET name=$1, drug_id=$2, dose=$3, date=$4 WHERE id=$5`
	stmt, err := r.db.PrepareContext(ctx, q)
	if err != nil {
		return Vaccination{}, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx,
		d.Name,
		d.DrugID,
		d.Dose,
		d.Date,
		id,
	)
	if err != nil {
		return Vaccination{}, err
	}
	ra, err := result.RowsAffected()
	if err != nil {
		return Vaccination{}, err
	} else if ra != 1 {
		return Vaccination{}, errors.New("expected 1 row affected")
	}
	d.ID = id
	return d, nil
}

func (r repository) DeleteVaccination(ctx context.Context, id int) error {
	q := `DELETE FROM vaccination WHERE id = $1`
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
