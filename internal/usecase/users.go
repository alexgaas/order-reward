package usecase

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
	repository "github.com/alexgaas/order-reward/internal/repo"
	"github.com/alexgaas/order-reward/internal/usecase/auth"
)

type UsersUseCase struct {
	repo Repository
}

// New -.
func New(r Repository) *UsersUseCase {
	return &UsersUseCase{
		repo: r,
	}
}

func (uc *UsersUseCase) RegisterUser(ctx context.Context, user domain.User) (string, error) {
	// save only hash of password in database
	auth.HashPassword(&user)

	// create user
	err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	// validate user in
	dbUser, err := uc.repo.GetUser(ctx, user.Login)
	if err != nil {
		return "", err
	}
	if dbUser.Password != user.Password {
		return "", repository.ErrInvalidLoginPassword
	}

	return auth.GetToken(user)
}
