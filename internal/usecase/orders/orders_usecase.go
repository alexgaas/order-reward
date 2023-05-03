package orders_usecase

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=orders_usecase

type (
	// Repository -.
	Repository interface {
		CreateOrder(ctx context.Context, login string, order domain.Order) error
		GetOrders(ctx context.Context, login string) ([]domain.Order, error)
	}
)