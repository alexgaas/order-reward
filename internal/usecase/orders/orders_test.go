package orders_usecase_test

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
	repository "github.com/alexgaas/order-reward/internal/repo"
	orders "github.com/alexgaas/order-reward/internal/usecase/orders"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type test struct {
	name string
	mock func()
	res  []domain.Order
	err  error
}

func MakeOrder(t *testing.T) (*orders.OrdersUseCase, *orders.MockRepository) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := orders.NewMockRepository(mockCtl)

	orderCase := orders.New(repo)

	return orderCase, repo
}

func MockOrders() []domain.Order {
	return []domain.Order{
		{
			ID:         1,
			UserID:     1,
			Number:     "10",
			Status:     "NEW",
			Accrual:    500.0,
			UploadedAt: time.Now().Unix(),
		},
		{
			ID:         2,
			UserID:     1,
			Number:     "20",
			Status:     "NEW",
			Accrual:    700.0,
			UploadedAt: time.Now().Unix(),
		},
	}
}

func TestGetOrders(t *testing.T) {
	t.Parallel()

	makeOrder, repo := MakeOrder(t)

	testLogin := "test_login"

	mockedOrders := MockOrders()

	tests := []test{
		{
			name: "orders not found",
			mock: func() {
				repo.EXPECT().GetOrders(context.Background(), testLogin).Return(nil, repository.ErrNoOrders)
			},
			res: nil,
			err: repository.ErrNoOrders,
		},
		{
			name: "user has been created successfully",
			mock: func() {
				repo.EXPECT().GetOrders(context.Background(), testLogin).Return(mockedOrders, nil)
			},
			res: mockedOrders,
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := makeOrder.GetOrders(context.Background(), testLogin)

			require.Exactly(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestMapOrdersToOrderResponse(t *testing.T) {
	mockedOrders := MockOrders()

	orderResponses := orders.MapOrdersToOrderResponse(mockedOrders)

	require.Exactly(t, orderResponses[0].Number, mockedOrders[0].Number)
	require.Exactly(t, orderResponses[1].Number, mockedOrders[1].Number)

	require.Exactly(t, orderResponses[0].Status, mockedOrders[0].Status)
	require.Exactly(t, orderResponses[1].Status, mockedOrders[1].Status)

	require.Exactly(t, orderResponses[0].Accrual, mockedOrders[0].Accrual)
	require.Exactly(t, orderResponses[1].Accrual, mockedOrders[1].Accrual)

	require.Exactly(t, orderResponses[0].UploadedAt,
		time.Unix(mockedOrders[0].UploadedAt, 0).Format(time.RFC3339))
	require.Exactly(t, orderResponses[0].UploadedAt,
		time.Unix(mockedOrders[1].UploadedAt, 0).Format(time.RFC3339))
}
