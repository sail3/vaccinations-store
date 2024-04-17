package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sail3/interfell-vaccinations/pkg/boilerplate"
)

func NewHTTPRouter(b boilerplate.Handler) http.Handler {

	r := chi.NewRouter()

	r.Get("/", b.Index)
	r.Get("/ping/{name}", b.Ping)
	r.Get("/message", b.ConsumeRepository)
	return r
}
