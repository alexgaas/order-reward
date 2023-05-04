package api

import (
	"encoding/json"
	"errors"
	"github.com/alexgaas/order-reward/internal/domain"
	repository "github.com/alexgaas/order-reward/internal/repo"
	users "github.com/alexgaas/order-reward/internal/usecase/users"
	"io"
	"net/http"
)

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

	token, err := users.New(app.Storage).RegisterUser(r.Context(), user)
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

func (app *AppHandler) Login(rw http.ResponseWriter, r *http.Request) {
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

	token, err := users.New(app.Storage).LoginUser(r.Context(), user, true)

	rw.Header().Set("Authorization", token)
	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write([]byte(""))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
