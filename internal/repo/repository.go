package repository

import (
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
		domain.Order{},
		domain.OrderLog{},
		domain.Account{},
	)
}
