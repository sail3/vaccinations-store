package boilerplate

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sail3/interfell-vaccinations/internal/response"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) Handler {
	return Handler{
		service: s,
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	response.ResponsdWithData(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "Boilerplate",
	})
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "name")
	response.ResponsdWithData(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("Pong %s", param),
	})
}

func (h *Handler) ConsumeRepository(w http.ResponseWriter, r *http.Request) {
	response.ResponsdWithData(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: h.service.GetMessage(),
	})
}
