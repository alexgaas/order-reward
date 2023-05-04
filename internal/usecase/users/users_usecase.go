package users_usecase

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
)

//go:generate mockgen -source=users_usecase.go -destination=./mocks_test.go -package=users_usecase

type (
	// Repository -.
	Repository interface {
		CreateUser(context.Context, domain.User) error
		GetUser(context.Context, string) (*domain.User, error)
	}
)
