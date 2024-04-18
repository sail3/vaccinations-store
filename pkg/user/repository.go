package user

import (
	"context"
	"database/sql"
	"time"
)

type Repository interface {
	RegisterUser(context.Context, User) (int, error)
	FindUser(context.Context, string) (User, error)
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

func (r repository) RegisterUser(ctx context.Context, d User) (int, error) {
	q := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
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
		d.Email,
		d.Password,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r repository) FindUser(ctx context.Context, e string) (User, error) {
	q := `SELECT id, name, email, password FROM users WHERE email = $1`
	rows := r.db.QueryRow(q, e)

	var d User
	rows.Scan(
		&d.ID,
		&d.Name,
		&d.Email,
		&d.Password,
	)
	return d, nil
}
