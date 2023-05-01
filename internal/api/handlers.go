package api

import (
	"encoding/json"
	"errors"
	"github.com/alexgaas/order-reward/internal/config"
	"github.com/alexgaas/order-reward/internal/domain"
	repository "github.com/alexgaas/order-reward/internal/repo"
	"github.com/alexgaas/order-reward/internal/usecase"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
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

	router.Group(func(r chi.Router) {
		r.Post("/api/user/register", app.Register)
	})

	return router
}

func (app *AppHandler) Register(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	user := domain.User{}

	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Login == "" || user.Password == "" {
		http.Error(rw, "wrong body format", http.StatusBadRequest)
		return
	}

	token, err := usecase.New(app.Storage).RegisterUser(r.Context(), user)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidLoginPassword) {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		} else if errors.Is(err, repository.ErrUserAlreadyExists) {
			http.Error(rw, err.Error(), http.StatusConflict)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Authorization", token)
	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write([]byte(""))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
