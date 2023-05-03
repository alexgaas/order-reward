package api

import (
	"encoding/json"
	"errors"
	repository "github.com/alexgaas/order-reward/internal/repo"
	orders "github.com/alexgaas/order-reward/internal/usecase/orders"
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
