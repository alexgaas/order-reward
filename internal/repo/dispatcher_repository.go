package repository

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
	"gorm.io/gorm"
)

func (db *Repository) DispatchGetOrders(ctx context.Context, status string) ([]string, error) {
	dbDispatchGetOrders := db.DB.WithContext(ctx)
	numList := make([]string, 0)

	rows, err := dbDispatchGetOrders.Model(&domain.Order{}).Select("number").Where("status = ?", status).Rows()
	defer rows.Close()

	if err != nil {
		return numList, err
	}

	for rows.Next() {
		var number string
		if err = rows.Scan(&number); err != nil {
			return numList, err
		}
		numList = append(numList, number)
	}
	if err = rows.Err(); err != nil {
		return numList, err
	}
	return numList, nil
}

func (db *Repository) DispatchUpdateOrder(ctx context.Context, order domain.Order) error {
	dbDispatchUpdateOrder := db.DB.WithContext(ctx)
	// transaction start
	return dbDispatchUpdateOrder.Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Model(&domain.Order{}).Where(
				"number = ?", order.Number,
			).Updates(
				domain.Order{Status: order.Status, Accrual: order.Accrual},
			).Error; err != nil {
				return err
			}
			if order.Accrual > 0 {
				if err := tx.Model(&domain.Account{}).Where(
					"user_id = (?)",
					tx.Model(&domain.Order{}).Select("id").Where(
						"number = ?", order.Number,
					),
				).UpdateColumn(
					"balance",
					gorm.Expr("balance + ?", order.Accrual),
				).Error; err != nil {
					return err
				}
			}
			return nil
		},
	)
	// transaction end
}
