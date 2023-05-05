package balance_usecase

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
)

type BalanceUseCase struct {
	repo Repository
}

func New(r Repository) *BalanceUseCase {
	return &BalanceUseCase{
		repo: r,
	}
}

func (uc *BalanceUseCase) GetBalance(ctx context.Context, login string) (*domain.Balance, error) {
	balance, err := uc.repo.GetBalance(ctx, login)

	if err != nil {
		return nil, err
	}

	return balance, err
}
