package user_test

import (
	"context"
	"os/user"
)

// Defined mocks to run tests
type mockRepository struct{}

func (m mockRepository) RegisterUser(context.Context, user.User) (int, error) {
	panic("implement me!!!")
}
func (m mockRepository) FindUser(context.Context, string) (user.User, error) {
	panic("implement me!!!")
}
