package usecase

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
	repository "github.com/alexgaas/order-reward/internal/repo"
	"github.com/alexgaas/order-reward/internal/usecase/auth"
)

func CreateUser(ctx context.Context, storage *repository.Repository, user domain.User) (string, error) {
	// save only hash of password in database
	auth.HashPassword(&user)

	// create user
	err := storage.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	// validate user in
	dbUser, err := storage.GetUser(ctx, user.Login)
	if err != nil {
		return "", err
	}
	if dbUser.Password != user.Password {
		return "", repository.ErrInvalidLoginPassword
	}

	return auth.GetToken(user)
}
