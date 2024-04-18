package drug_test

import (
	"context"

	"github.com/sail3/interfell-vaccinations/pkg/drug"
)

// Defined mocks to run tests
type mockRepository struct{}

func (m mockRepository) RegisterDrug(context.Context, drug.Drug) (int, error) {
	panic("implement me!!")
}
func (m mockRepository) ListDrug(context.Context) ([]drug.Drug, error) {
	panic("implement me!!")
}
func (m mockRepository) UpdateDrug(context.Context, int, drug.Drug) (drug.Drug, error) {
	panic("implement me!!")
}
func (m mockRepository) DeleteDrug(context.Context, int) error {
	panic("implement me!!")
}
