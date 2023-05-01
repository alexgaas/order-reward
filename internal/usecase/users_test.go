package usecase_test

import (
	"context"
	"github.com/alexgaas/order-reward/internal/domain"
	repository "github.com/alexgaas/order-reward/internal/repo"
	"github.com/alexgaas/order-reward/internal/usecase"
	"github.com/alexgaas/order-reward/internal/usecase/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func CreateUser(t *testing.T) (*usecase.UsersUseCase, *MockRepository) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockRepository(mockCtl)

	userCase := usecase.New(repo)

	return userCase, repo
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	createUser, repo := CreateUser(t)

	// user model for repo call
	testRepoUser := domain.User{
		Login:    "test_user",
		Password: "test_password",
	}
	// user model for use case call
	testUsecaseUser := testRepoUser

	// user model for repo call must have hidden password already
	auth.HashPassword(&testRepoUser)

	tokenToValidate, _ := auth.GetToken(testUsecaseUser)

	tests := []test{
		{
			name: "user with such credentials already exist",
			mock: func() {
				repo.EXPECT().CreateUser(context.Background(), testRepoUser).Return(repository.ErrUserAlreadyExists)
			},
			res: "",
			err: repository.ErrUserAlreadyExists,
		},
		{
			name: "invalid login/password",
			mock: func() {
				repo.EXPECT().CreateUser(context.Background(), testRepoUser).Return(repository.ErrInvalidLoginPassword)
			},
			res: "",
			err: repository.ErrInvalidLoginPassword,
		},
		{
			name: "user has been created successfully",
			mock: func() {
				repo.EXPECT().GetUser(context.Background(), testRepoUser.Login).Return(&testRepoUser, nil)
				repo.EXPECT().CreateUser(context.Background(), testRepoUser).Return(nil)
			},
			res: tokenToValidate,
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := createUser.RegisterUser(context.Background(), testUsecaseUser)

			require.Exactly(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}