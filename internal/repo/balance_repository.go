package repository

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
)

func (db *Repository) GetBalance(ctx context.Context, login string) (*domain.Balance, error) {
	dbBalance := db.DB.WithContext(ctx)

	var result domain.Balance

	err := dbBalance.Model(&domain.User{}).Select(
		"sum(order_logs.sum) as summary, accounts.balance as balance",
	).Joins(
		"left join order_logs on order_logs.user_id = users.id",
	).Joins(
		"left join accounts on accounts.user_id = users.id",
	).Group("accounts.balance").Where("users.login = ?", login).Scan(&result).Error

	return &result, err
}
