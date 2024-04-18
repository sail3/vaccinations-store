package vaccination

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var d VaccinationRequest

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		_ = response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println(d)
	res, err := h.service.RegisterVaccination(ctx, d)
	if err != nil {
		_ = response.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponsdWithData(w, http.StatusOK, res)
}

func (h *Handler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println(ctx)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		_ = response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	var d VaccinationRequest
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		_ = response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.service.UpdateVaccination(ctx, id, d)
	if err != nil {
		_ = response.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponsdWithData(w, http.StatusOK, res)

}

func (h *Handler) ListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := h.service.ListVaccination(ctx)
	if err != nil {
		_ = response.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}
	response.ResponsdWithData(w, http.StatusOK, res)
}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println(ctx)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		_ = response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.DeleteVaccination(ctx, id)
	if err != nil {
		_ = response.ResponseWithError(w, http.StatusNotFound, err)
		return
	}
	response.ResponsdWithData(w, http.StatusNoContent, "")
}
