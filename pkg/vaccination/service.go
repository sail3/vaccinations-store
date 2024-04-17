package vaccination

import (
	"context"
	"time"
)

type Service interface {
	RegisterVaccination(context.Context, VaccinationRequest) (Vaccination, error)
	UpdateVaccination(context.Context, int, VaccinationRequest) (Vaccination, error)
	ListVaccination(context.Context) ([]Vaccination, error)
	DeleteVaccination(context.Context, int) error
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

type service struct {
	repository Repository
}

func (s service) RegisterVaccination(ctx context.Context, dr VaccinationRequest) (Vaccination, error) {
	vac := Vaccination{
		Name:   dr.Name,
		DrugID: dr.DrugID,
		Dose:   dr.Dose,
		Date:   time.Time(dr.Date),
	}
	id, err := s.repository.RegisterVaccination(ctx, vac)
	if err != nil {
		return Vaccination{}, err
	}
	vac.ID = id
	return vac, nil
}

func (s service) UpdateVaccination(ctx context.Context, id int, dr VaccinationRequest) (Vaccination, error) {
	vac := Vaccination{
		Name:   dr.Name,
		DrugID: dr.DrugID,
		Dose:   dr.Dose,
		Date:   time.Time(dr.Date),
	}

	res, err := s.repository.UpdateVaccination(ctx, id, vac)
	if err != nil {
		return Vaccination{}, err
	}
	res.ID = id
	return res, nil
}

func (s service) ListVaccination(ctx context.Context) ([]Vaccination, error) {
	res, err := s.repository.ListVaccination(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) DeleteVaccination(ctx context.Context, id int) error {
	err := s.repository.DeleteVaccination(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
