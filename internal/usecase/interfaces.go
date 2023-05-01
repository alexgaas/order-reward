package usecase

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Repository -.
	Repository interface {
		CreateUser(context.Context, domain.User) error
		GetUser(context.Context, string) (*domain.User, error)
	}
)
