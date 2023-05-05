package repository

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
	"gorm.io/gorm"
)

func (db *Repository) GetOrders(ctx context.Context, login string) ([]domain.Order, error) {
	dbOrders := db.DB.WithContext(ctx)
	orders := make([]domain.Order, 0)
	user, err := db.GetUser(ctx, login)
	if err != nil {
		return orders, err
	}

	dbOrders.Where("user_id = ?", user.ID).Find(&orders)
	if len(orders) == 0 {
		return orders, ErrNoOrders
	}
	return orders, err
}

func (db *Repository) SaveOrder(ctx context.Context, login string, order domain.Order) error {
	dbOrders := db.DB.WithContext(ctx)
	user, err := db.GetUser(ctx, login)
	if err != nil {
		return err
	}
	order.UserID = user.ID

	var dbOrder domain.Order
	dbOrders.Where("number = ?", order.Number).Find(&dbOrder)
	if dbOrder.Number != "" {
		if dbOrder.UserID == user.ID {
			return ErrOrderExists
		} else {
			return ErrOrderExistsAnother
		}
	}
	return dbOrders.Create(&order).Error
}

func (db *Repository) GetOrderLog(ctx context.Context, login string) ([]domain.OrderLog, error) {
	dbOrderList := db.DB.WithContext(ctx)
	orders := make([]domain.OrderLog, 0)
	user, err := db.GetUser(ctx, login)
	if err != nil {
		return orders, err
	}

	err = dbOrderList.Where("user_id = ?", user.ID).Find(&orders).Error
	if len(orders) == 0 {
		return orders, ErrNoOrders
	}
	return orders, err
}

func (db *Repository) WithdrawOrder(ctx context.Context, login string, orderLog domain.OrderLog) error {
	dbWithdrawOrder := db.DB.WithContext(ctx)

	// transaction start
	return dbWithdrawOrder.Transaction(
		func(tx *gorm.DB) error {
			var user domain.User
			if err := dbWithdrawOrder.First(&user, "login = ?", login).Error; err != nil {
				return err
			}
			// compare balance to sum
			balance := domain.Account{UserID: user.ID}
			if err := tx.Select("balance").Find(&balance).Error; err != nil {
				return err
			}
			if balance.Balance.Float64 < orderLog.Sum {
				return ErrNotEnoughFunds
			}
			// balance - sum
			if err := tx.Model(&domain.Account{}).Where(
				"user_id = (?)", user.ID,
			).UpdateColumn("balance", gorm.Expr("balance - ?", orderLog.Sum)).Error; err != nil {
				return err
			}
			// write orderLog entry
			orderLog.UserID = user.ID
			return tx.Create(&orderLog).Error
		},
	)
	// transaction end
}
