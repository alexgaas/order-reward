package orders_usecase

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
	"time"
)

type OrdersUseCase struct {
	repo Repository
}

func New(r Repository) *OrdersUseCase {
	return &OrdersUseCase{
		repo: r,
	}
}

func (uc *OrdersUseCase) GetOrders(ctx context.Context, login string) ([]domain.Order, error) {
	orders, err := uc.repo.GetOrders(ctx, login)

	if err != nil {
		return nil, err
	}

	return orders, err
}

func MapOrdersToOrderResponse(orders []domain.Order) []domain.OrderResponse {
	orderResp := make([]domain.OrderResponse, 0)
	for _, order := range orders {
		resp := domain.OrderResponse{
			Number:     order.Number,
			Status:     order.Status,
			UploadedAt: time.Unix(order.UploadedAt, 0).Format(time.RFC3339),
		}
		if order.Accrual > 0 {
			resp.Accrual = order.Accrual
		}
		orderResp = append(orderResp, resp)
	}

	return orderResp
}
