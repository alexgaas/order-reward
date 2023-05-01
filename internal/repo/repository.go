package repository

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// New -.
func New(conn *gorm.DB) *Repository {
	return &Repository{DB: conn}
}

func NewDB(dsn string) (Repository, error) {
	conn, err := gorm.Open(
		sqlite.Open(dsn),
		&gorm.Config{},
	)

	return *New(conn), err
}

func (db *Repository) CloseDB() error {
	closeDb, err := db.DB.DB()
	closeErr := closeDb.Close()
	if closeErr != nil {
		return closeErr
	}
	return err
}

func (db *Repository) InitDB() error {
	return db.DB.AutoMigrate(
		domain.User{},
	)
}

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
