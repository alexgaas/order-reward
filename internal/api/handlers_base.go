package api

import (
	"github.com/alexgaas/order-reward/internal/config"
	repository "github.com/alexgaas/order-reward/internal/repo"
	"github.com/go-chi/chi/v5"
	"log"
)

type AppHandler struct {
	AccrualAddress string
	Storage        *repository.Repository
	Logger         *log.Logger
}

func NewAppHandler(config config.Config, logger *log.Logger) (*AppHandler, error) {
	app := &AppHandler{
		AccrualAddress: config.AccrualAddress,
		Logger:         logger,
	}

	dbContext, err := repository.NewDB(config.DatabaseDSN)
	if err != nil {
		return app, err
	}
	app.Storage = &dbContext

	return app, nil
}

func NewRouter(app *AppHandler) *chi.Mux {
	router := chi.NewRouter()
	router.Use(GzipMiddle)

	router.Group(func(r chi.Router) {
		r.Post("/api/user/register", app.Register)
		r.Post("/api/user/login", app.Login)
	})

	router.Group(func(r chi.Router) {
		r.Use(Authenticator)
		r.Get("/api/orders", app.GetOrders)
		r.Post("/api/orders", app.PostOrder)
		r.Post("/api/orders/withdraw", app.Withdraw)
		r.Post("/api/orders/withdrawals", app.Withdrawals)
		r.Get("/api/balance", app.GetBalance)
	})

	return router
}
