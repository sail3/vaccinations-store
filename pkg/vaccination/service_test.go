package vaccination_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"github.com/sail3/interfell-vaccinations/pkg/vaccination"
)

type registerVaccinationMockRepository struct {
	mockRepository
	registerVaccinationResp int
	error                   error
}

func (m registerVaccinationMockRepository) RegisterVaccination(c context.Context, v vaccination.Vaccination) (int, error) {
	return m.registerVaccinationResp, m.error
}

func TestService_RegisterVaccination(t *testing.T) {
	currentTime := time.Now()
	someError := errors.New("some error")
	vaccinationRequestMock := vaccination.VaccinationRequest{
		Name:   "test",
		DrugID: 12,
		Dose:   10,
		Date:   vaccination.CustomTime(currentTime),
	}
	vaccinationMock := vaccination.Vaccination{
		ID:     1,
		Name:   "test",
		DrugID: 12,
		Dose:   10,
		Date:   currentTime,
	}
	tests := []struct {
		name                    string
		inputVacination         vaccination.VaccinationRequest
		RegisterVaccinationResp int
		RegisterVaccinationErr  error
		resp                    vaccination.Vaccination
		err                     error
	}{
		{
			name:                    "Sucess",
			inputVacination:         vaccinationRequestMock,
			RegisterVaccinationResp: 1,
			RegisterVaccinationErr:  nil,
			resp:                    vaccinationMock,
			err:                     nil,
		},
		{
			name:                    "Error",
			inputVacination:         vaccinationRequestMock,
			RegisterVaccinationResp: 0,
			RegisterVaccinationErr:  someError,
			resp:                    vaccination.Vaccination{},
			err:                     someError,
		},
	}
	for _, test := range tests {
		m := registerVaccinationMockRepository{
			registerVaccinationResp: test.RegisterVaccinationResp,
			error:                   test.RegisterVaccinationErr,
		}

		s := vaccination.NewService(m)
		ctx := context.Background()

		resp, err := s.RegisterVaccination(ctx, test.inputVacination)
		assert.Equal(t, test.resp, resp)
		assert.Equal(t, test.err, err)
	}
}

type updateVaccinationMockRepository struct {
	mockRepository
	resp vaccination.Vaccination
	err  error
}

func (m updateVaccinationMockRepository) UpdateVaccination(c context.Context, id int, v vaccination.Vaccination) (vaccination.Vaccination, error) {
	return m.resp, m.err
}

func TestService_UpdateVaccination(t *testing.T) {
	currentTime := time.Now()
	someError := errors.New("some error")
	vaccinationRequestMock := vaccination.VaccinationRequest{
		Name:   "test",
		DrugID: 12,
		Dose:   10,
		Date:   vaccination.CustomTime(currentTime),
	}
	vaccinationMock := vaccination.Vaccination{
		ID:     1,
		Name:   "test",
		DrugID: 12,
		Dose:   10,
		Date:   currentTime,
	}
	tests := []struct {
		name                             string
		inputVacination                  vaccination.VaccinationRequest
		inputID                          int
		updateVaccinationRepoVaccination vaccination.Vaccination
		updateVaccinationErr             error
		resp                             vaccination.Vaccination
		err                              error
	}{
		{
			name:                             "Sucess",
			inputVacination:                  vaccinationRequestMock,
			inputID:                          1,
			updateVaccinationRepoVaccination: vaccinationMock,
			updateVaccinationErr:             nil,
			resp:                             vaccinationMock,
			err:                              nil,
		},
		{
			name:                             "Fail",
			inputVacination:                  vaccinationRequestMock,
			inputID:                          1,
			updateVaccinationRepoVaccination: vaccinationMock,
			updateVaccinationErr:             someError,
			resp:                             vaccination.Vaccination{},
			err:                              someError,
		},
	}
	for _, test := range tests {
		m := updateVaccinationMockRepository{
			resp: test.updateVaccinationRepoVaccination,
			err:  test.updateVaccinationErr,
		}

		s := vaccination.NewService(m)
		ctx := context.Background()

		resp, err := s.UpdateVaccination(ctx, test.inputID, test.inputVacination)
		assert.Equal(t, test.resp, resp)
		assert.Equal(t, test.err, err)
	}
}

type ListVaccinationRepositoryMock struct {
	mockRepository
	resp []vaccination.Vaccination
	err  error
}

func (m ListVaccinationRepositoryMock) ListVaccination(context.Context) ([]vaccination.Vaccination, error) {
	return m.resp, m.err
}

func TestService_ListVaccination(t *testing.T) {
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
		name                           string
		listVaccinationMockVaccination []vaccination.Vaccination
		listVaccinationMockErr         error
		resp                           []vaccination.Vaccination
		err                            error
	}{
		{
			name:                           "Sucess",
			listVaccinationMockVaccination: arrVaccination,
			listVaccinationMockErr:         nil,
			resp:                           arrVaccination,
			err:                            nil,
		},
		{
			name:                           "Fail!!",
			listVaccinationMockVaccination: nil,
			listVaccinationMockErr:         someError,
			resp:                           nil,
			err:                            someError,
		},
	}
	for _, test := range tests {
		m := ListVaccinationRepositoryMock{
			resp: test.listVaccinationMockVaccination,
			err:  test.listVaccinationMockErr,
		}

		s := vaccination.NewService(m)
		ctx := context.Background()

		resp, err := s.ListVaccination(ctx)
		assert.Equal(t, test.resp, resp)
		assert.Equal(t, test.err, err)
	}
}

type DeleteVaccinationRepositoryMock struct {
	mockRepository
	err error
}

func (m DeleteVaccinationRepositoryMock) DeleteVaccination(context.Context, int) error {
	return m.err
}

func TestService_DeleteVaccination(t *testing.T) {
	someError := errors.New("some error")

	tests := []struct {
		name                     string
		inputID                  int
		deleteVaccinationMockErr error
		err                      error
	}{
		{
			name:                     "Sucess",
			inputID:                  10,
			deleteVaccinationMockErr: nil,
			err:                      nil,
		},
		{
			name:                     "Fail!!",
			inputID:                  10,
			deleteVaccinationMockErr: someError,
			err:                      someError,
		},
	}
	for _, test := range tests {
		m := DeleteVaccinationRepositoryMock{
			err: test.deleteVaccinationMockErr,
		}

		s := vaccination.NewService(m)
		ctx := context.Background()

		err := s.DeleteVaccination(ctx, test.inputID)
		assert.Equal(t, test.err, err)
	}
}

// Defined Mocks to run tests
type mockRepository struct{}

func (m mockRepository) RegisterVaccination(context.Context, vaccination.Vaccination) (int, error) {
	panic("implement me!!")
}
func (m mockRepository) ListVaccination(context.Context) ([]vaccination.Vaccination, error) {
	panic("implement me!!")
}
func (m mockRepository) UpdateVaccination(context.Context, int, vaccination.Vaccination) (vaccination.Vaccination, error) {
	panic("implement me!!")
}
func (m mockRepository) DeleteVaccination(context.Context, int) error {
	panic("implement me!!")
}
