package api

import (
	"encoding/json"
	"errors"
	repository "github.com/alexgaas/order-reward/internal/repo"
	orders "github.com/alexgaas/order-reward/internal/usecase/orders"
	"io"
	"net/http"
)

func (app *AppHandler) GetOrders(rw http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("Login")

	rw.Header().Set("Content-Type", "application/json")

	getOrders, err := orders.New(app.Storage).GetOrders(r.Context(), login)
	if err != nil {
		if errors.Is(err, repository.ErrNoOrders) {
			http.Error(rw, err.Error(), http.StatusNoContent)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(orders.MapOrdersToOrderResponse(getOrders))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(resp)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *AppHandler) PostOrder(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	bodyStr := string(body)

	login := r.Header.Get("Login")

	contentType := r.Header.Get("Content-type")
	if contentType != "text/plain" || bodyStr == "" {
		http.Error(rw, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := orders.New(app.Storage).CreateOrder(r.Context(), login, bodyStr); err != nil {
		if errors.Is(err, repository.ErrOrderExists) {
			http.Error(rw, err.Error(), http.StatusOK)
			return
		}
		if errors.Is(err, repository.ErrOrderExistsAnother) {
			http.Error(rw, err.Error(), http.StatusConflict)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusAccepted)
	_, err = rw.Write([]byte(""))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
