package repository

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
)

func (db *Repository) CreateUser(ctx context.Context, user domain.User) error {
	dbContext := db.DB.WithContext(ctx)
	var exists bool
	dbContext.Model(&domain.User{}).
		Select("count(*) > 0").
		Where("login = ?", user.Login).
		Find(&exists)
	if exists {
		return ErrUserAlreadyExists
	}

	return dbContext.Create(&user).Error
}

func (db *Repository) GetUser(ctx context.Context, login string) (*domain.User, error) {
	var user domain.User
	dbContext := db.DB.WithContext(ctx)
	err := dbContext.First(&user, "login = ?", login).Error
	return &user, err
}
