package api

import (
	"encoding/json"
	balance "github.com/alexgaas/order-reward/internal/usecase/balance"
	"net/http"
)

func (app *AppHandler) GetBalance(rw http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("Login")

	result, err := balance.New(app.Storage).GetBalance(r.Context(), login)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(resultJSON)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
