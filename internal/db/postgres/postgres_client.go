package postgres

import (
	"context"
	"database/sql"
	"sync"
	"time"
)

var (
	postgresClient *Client
	postgresOnce   sync.Once
)

type Client struct {
	*sql.DB
}

func NewPostgresClient(source string) *Client {
	postgresOnce.Do(func() {
		db, err := sql.Open("postgres", source)
		if err != nil {
			panic(err)
		}

		err = db.PingContext(context.Background())
		if err != nil {
			panic(err)
		}

		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(2 * time.Minute)

		postgresClient = &Client{db}
	})

	return postgresClient
}
