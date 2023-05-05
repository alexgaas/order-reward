package integration_test

import (
	. "github.com/Eun/go-hit"
	base "github.com/alexgaas/order-reward/integration-test"

	"net/http"
	"testing"
)

func TestGetBalance_NoBalance(t *testing.T) {
	jwt := RegisterUser()

	// test orders not found
	Test(t,
		Description("Get the reward balance for the user"),
		Get(base.BasePath+"balance"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Headers("Authorization").Add(jwt),
		Expect().Status().Equal(http.StatusInternalServerError),
	)
}
