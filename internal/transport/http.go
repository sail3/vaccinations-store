package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sail3/interfell-vaccinations/pkg/drug"
	"github.com/sail3/interfell-vaccinations/pkg/vaccination"
)

func NewHTTPRouter(d drug.Handler, v vaccination.Handler) http.Handler {

	r := chi.NewRouter()

	r.Post("/drugs", d.RegisterHandler)
	r.Put("/drugs/{id}", d.UpdateHandler)
	r.Get("/drugs", d.ListHandler)
	r.Delete("/drugs/{id}", d.DeleteHandler)

	r.Post("/vaccination", v.RegisterHandler)
	r.Put("/vaccination/{id}", v.UpdateHandler)
	r.Get("/vaccination", v.ListHandler)
	r.Delete("/vaccination/{id}", v.DeleteHandler)

	return r
}
