package drug

import (
	"context"
	"fmt"
	"time"
)

type Service interface {
	RegisterDrug(context.Context, RegisterDrugRequest) (Drug, error)
	UpdateDrug(context.Context, int, UpdateDrugRequest) (Drug, error)
	ListDrug(context.Context) ([]Drug, error)
	DeleteDrug(context.Context, int) error
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

type service struct {
	repository Repository
}

func (s service) RegisterDrug(ctx context.Context, dr RegisterDrugRequest) (Drug, error) {
	drug := Drug{
		Name:        dr.Name,
		Approved:    dr.Approved,
		MinDose:     dr.MinDose,
		MaxDose:     dr.MaxDose,
		AvailableAt: time.Time(dr.AvailableAt),
	}
	id, err := s.repository.RegisterDrug(ctx, drug)
	fmt.Println(err)
	if err != nil {
		return Drug{}, err
	}
	drug.ID = id
	return drug, nil
}

func (s service) UpdateDrug(ctx context.Context, id int, dr UpdateDrugRequest) (Drug, error) {
	drug := Drug{
		Name:        dr.Name,
		Approved:    dr.Approved,
		MinDose:     dr.MinDose,
		MaxDose:     dr.MaxDose,
		AvailableAt: time.Time(dr.AvailableAt),
	}

	res, err := s.repository.UpdateDrug(ctx, id, drug)
	if err != nil {
		return Drug{}, err
	}
	res.ID = id
	return res, nil
}

func (s service) ListDrug(ctx context.Context) ([]Drug, error) {
	res, err := s.repository.ListDrug(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s service) DeleteDrug(ctx context.Context, id int) error {
	err := s.repository.DeleteDrug(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
