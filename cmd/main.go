package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sail3/interfell-vaccinations/internal/transport"
	"github.com/sail3/interfell-vaccinations/pkg/boilerplate"
)

func main() {

	br := boilerplate.NewRepository()
	bs := boilerplate.NewService(br)
	bh := boilerplate.NewHandler(bs)

	r := transport.NewHTTPRouter(bh)
	srv := http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	srv.ListenAndServe()
}

func router() http.Handler {
	r := chi.NewRouter()

	return r
}
