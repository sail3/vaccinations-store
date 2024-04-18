package user

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (h Handler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var s SignupRequest

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		_ = response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Println(s)

	user, err := h.service.SignupService(ctx, s)
	if err != nil {
		_ = response.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponsdWithData(w, http.StatusOK, user)
}

func (h Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var lr LoginRequest

	err := json.NewDecoder(r.Body).Decode(&lr)
	if err != nil {
		_ = response.ResponseWithError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.service.LoginService(ctx, lr)
	if err != nil {
		_ = response.ResponseWithError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponsdWithData(w, http.StatusOK, res)

}
