package balance_usecase

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
)

//go:generate mockgen -source=balance_usecase.go -destination=./mocks_test.go -package=balance_usecase

type (
	// Repository -.
	Repository interface {
		GetBalance(ctx context.Context, login string) (*domain.Balance, error)
	}
)
