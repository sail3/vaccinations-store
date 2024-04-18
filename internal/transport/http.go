package transport

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sail3/interfell-vaccinations/internal/transport/middleware"
	"github.com/sail3/interfell-vaccinations/pkg/drug"
	"github.com/sail3/interfell-vaccinations/pkg/user"
	"github.com/sail3/interfell-vaccinations/pkg/vaccination"
)

func NewHTTPRouter(st string, d drug.Handler, v vaccination.Handler, u user.Handler) http.Handler {

	r := chi.NewRouter()

	r.Post("/signup", u.SignupHandler)
	r.Post("/login", u.LoginHandler)

	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticator(st))
		r.Post("/drugs", d.RegisterHandler)
		r.Put("/drugs/{id}", d.UpdateHandler)
		r.Get("/drugs", d.ListHandler)
		r.Delete("/drugs/{id}", d.DeleteHandler)

		r.Post("/vaccination", v.RegisterHandler)
		r.Put("/vaccination/{id}", v.UpdateHandler)
		r.Get("/vaccination", v.ListHandler)
		r.Delete("/vaccination/{id}", v.DeleteHandler)
	})

	r.Handle("/swagger/*", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger"))))

	return r
}
