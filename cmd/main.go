package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/sail3/interfell-vaccinations/internal/config"
	"github.com/sail3/interfell-vaccinations/internal/db/postgres"
	"github.com/sail3/interfell-vaccinations/internal/transport"
	"github.com/sail3/interfell-vaccinations/pkg/drug"
	"github.com/sail3/interfell-vaccinations/pkg/user"
	"github.com/sail3/interfell-vaccinations/pkg/vaccination"
)

func main() {
	conf := config.New()

	err := doMigrate(conf.DbURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db := postgres.NewPostgresClient(conf.DbURL)

	dr := drug.NewRepository(db.DB)
	ds := drug.NewService(dr)
	dh := drug.NewHandler(ds)

	vr := vaccination.NewRepository(db.DB)
	vs := vaccination.NewService(vr)
	vh := vaccination.NewHandler(vs)

	ur := user.NewRepository(db.DB)
	us := user.NewService(ur, conf.SignatureToken)
	uh := user.NewHandler(us)

	r := transport.NewHTTPRouter(conf.SignatureToken, dh, vh, uh)
	srv := http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	srv.ListenAndServe()
}

const migrationsRootFolder = "file://migrations"

func doMigrate(databaseURL string) error {
	m, err := migrate.New(
		migrationsRootFolder,
		databaseURL,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}
