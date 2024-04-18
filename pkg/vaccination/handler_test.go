package vaccination_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/magiconair/properties/assert"
	"github.com/sail3/interfell-vaccinations/pkg/vaccination"
	"github.com/stretchr/testify/require"
)

type registerVaccinationServiceMock struct {
	mockService
	resp vaccination.Vaccination
	err  error
}

func (r registerVaccinationServiceMock) RegisterVaccination(context.Context, vaccination.VaccinationRequest) (vaccination.Vaccination, error) {
	return r.resp, r.err
}
func TestHandler_RegisterHandler(t *testing.T) {
	vt := vaccination.VaccinationRequest{
		Name:   "name test",
		DrugID: 10,
		Dose:   10,
		Date:   vaccination.CustomTime(time.Now()),
	}
	someError := errors.New("some error")
	tests := []struct {
		name            string
		mockResp        vaccination.Vaccination
		mockErr         error
		inputVacination vaccination.VaccinationRequest
		statusCode      int
		err             error
	}{
		{
			name:            "Success",
			mockResp:        vaccination.Vaccination{},
			mockErr:         nil,
			inputVacination: vt,
			statusCode:      http.StatusOK,
			err:             nil,
		},
		{
			name:            "Fail",
			mockResp:        vaccination.Vaccination{},
			mockErr:         someError,
			inputVacination: vt,
			statusCode:      http.StatusInternalServerError,
			err:             someError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := registerVaccinationServiceMock{
				resp: test.mockResp,
				err:  test.mockErr,
			}
			msh, _ := json.Marshal(test.inputVacination)
			r, err := http.NewRequest(http.MethodPost, "/vaccination", strings.NewReader(string(msh)))
			w := httptest.NewRecorder()
			if err != nil {
				require.NoError(t, err)
			}

			h := vaccination.NewHandler(m)

			mux := chi.NewMux()
			mux.Post("/vaccination", h.RegisterHandler)
			mux.ServeHTTP(w, r)

			statusCode := w.Result().StatusCode
			assert.Equal(t, test.statusCode, statusCode)

		})
	}
}

type updateVaccinationServiceMock struct {
	mockService
	resp vaccination.Vaccination
	err  error
}

func (r updateVaccinationServiceMock) UpdateVaccination(context.Context, int, vaccination.VaccinationRequest) (vaccination.Vaccination, error) {
	return r.resp, r.err
}

func TestHandler_UpdateHandler(t *testing.T) {
	vt := vaccination.VaccinationRequest{
		Name:   "name test",
		DrugID: 10,
		Dose:   10,
		Date:   vaccination.CustomTime(time.Now()),
	}
	tests := []struct {
		name            string
		mockResp        vaccination.Vaccination
		mockErr         error
		inputID         int
		inputVacination vaccination.VaccinationRequest
		statusCode      int
		err             error
	}{
		{
			name:            "Success!!",
			mockResp:        vaccination.Vaccination{},
			mockErr:         nil,
			inputID:         10,
			inputVacination: vt,
			statusCode:      http.StatusBadRequest,
			err:             nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := updateVaccinationServiceMock{
				resp: test.mockResp,
				err:  test.mockErr,
			}
			msh, _ := json.Marshal(test.inputVacination)
			r, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/vaccination/%d", test.inputID), strings.NewReader(string(msh)))
			w := httptest.NewRecorder()
			if err != nil {
				require.NoError(t, err)
			}

			h := vaccination.NewHandler(m)

			mux := chi.NewMux()
			mux.Put("/vaccination/{id}", h.UpdateHandler)
			mux.ServeHTTP(w, r)

			statusCode := w.Result().StatusCode
			assert.Equal(t, test.statusCode, statusCode)
		})
	}
}

type listVaccinationServiceMock struct {
	mockService
	resp []vaccination.Vaccination
	err  error
}

func (m listVaccinationServiceMock) ListVaccination(context.Context) ([]vaccination.Vaccination, error) {
	return m.resp, m.err
}

func TestHandler_ListHandler(t *testing.T) {
	currentTime := time.Now()
	someError := errors.New("some error")
	vaccinationMock := vaccination.Vaccination{
		ID:     1,
		Name:   "test",
		DrugID: 12,
		Dose:   10,
		Date:   currentTime,
	}
	arrVaccination := []vaccination.Vaccination{
		vaccinationMock, vaccinationMock,
	}
	tests := []struct {
		name       string
		mockResp   []vaccination.Vaccination
		mockErr    error
		statusCode int
		err        error
	}{
		{
			name:       "Success!!",
			mockResp:   arrVaccination,
			mockErr:    nil,
			statusCode: http.StatusOK,
			err:        nil,
		},
		{
			name:       "Fail!!",
			mockResp:   nil,
			mockErr:    someError,
			statusCode: http.StatusInternalServerError,
			err:        someError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := listVaccinationServiceMock{
				resp: test.mockResp,
				err:  test.mockErr,
			}

			r, err := http.NewRequest(http.MethodPut, "/vaccination", nil)
			w := httptest.NewRecorder()
			if err != nil {
				require.NoError(t, err)
			}

			h := vaccination.NewHandler(s)

			mux := chi.NewMux()
			mux.Put("/vaccination", h.ListHandler)
			mux.ServeHTTP(w, r)

			statusCode := w.Result().StatusCode
			assert.Equal(t, test.statusCode, statusCode)
		})
	}
}

type deleteVaccinationServiceMock struct {
	mockService
	err error
}

func (m deleteVaccinationServiceMock) DeleteVaccination(context.Context, int) error {
	return m.err
}

func TestHandler_DeleteHandler(t *testing.T) {
	tests := []struct {
		name       string
		mockErr    error
		inputID    int
		statusCode int
		err        error
	}{
		{
			name:       "Success!!",
			mockErr:    nil,
			inputID:    10,
			statusCode: http.StatusBadRequest,
			err:        nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := deleteVaccinationServiceMock{
				err: test.mockErr,
			}

			r, err := http.NewRequest(http.MethodDelete, "/vaccination/10", nil)
			w := httptest.NewRecorder()
			if err != nil {
				require.NoError(t, err)
			}

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", strconv.Itoa(test.inputID))
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			h := vaccination.NewHandler(s)

			mux := chi.NewMux()
			mux.Delete("/vaccination/{id}", h.DeleteHandler)
			mux.ServeHTTP(w, r)

			statusCode := w.Result().StatusCode
			assert.Equal(t, statusCode, test.statusCode)
		})
	}
}

type mockService struct{}

func (m mockService) RegisterVaccination(context.Context, vaccination.VaccinationRequest) (vaccination.Vaccination, error) {
	panic("Implement me!!")
}
func (m mockService) UpdateVaccination(context.Context, int, vaccination.VaccinationRequest) (vaccination.Vaccination, error) {
	panic("Implement me!!")
}
func (m mockService) ListVaccination(context.Context) ([]vaccination.Vaccination, error) {
	panic("Implement me!!")
}
func (m mockService) DeleteVaccination(context.Context, int) error {
	panic("Implement me!!")
}
