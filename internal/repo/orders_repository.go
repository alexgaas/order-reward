package repository

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
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
