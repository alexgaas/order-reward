package domain

type Balance struct {
	Balance float64 `json:"current"`
	Summary float64 `json:"withdrawn"`
}
