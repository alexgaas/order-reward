package balance_usecase_test

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
	balance "github.com/alexgaas/order-reward/internal/usecase/balance"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

type test struct {
	name string
	mock func()
	res  *domain.Balance
	err  error
	num  string
}

func MakeBalance(t *testing.T) (*balance.BalanceUseCase, *balance.MockRepository) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := balance.NewMockRepository(mockCtl)

	balanceCase := balance.New(repo)

	return balanceCase, repo
}

func TestGetOrders(t *testing.T) {
	t.Parallel()

	makeBalance, repo := MakeBalance(t)

	testLogin := "test_login"

	mockBalance := domain.Balance{
		Balance: 100.0,
		Summary: 500.0,
	}

	tests := []test{
		{
			name: "get orders successfully",
			mock: func() {
				repo.EXPECT().GetBalance(context.Background(), testLogin).Return(&mockBalance, nil)
			},
			res: &mockBalance,
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := makeBalance.GetBalance(context.Background(), testLogin)

			require.Exactly(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
