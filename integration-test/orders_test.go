package integration_test

import (
	"fmt"
	base "github.com/alexgaas/order-reward/integration-test"
	"github.com/google/uuid"
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
)

func TestGetOrders_OrdersNotFound(t *testing.T) {
	var jwt string

	// register random user to get JWT
	randomLogin := uuid.New()
	randomPassword := uuid.New()
	body := fmt.Sprintf(`{
		"login": "test_%s",
		"password": "test_%s"
	}`, randomLogin, randomPassword)
	MustDo(
		Post(base.BasePath+"user/register"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
		Store().Response().Headers("Authorization").In(&jwt),
	)

	// test orders not found
	Test(t,
		Description("Retrieve a list of user order numbers, including their processing status and accrual information"),
		Get(base.BasePath+"orders"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Headers("Authorization").Add(jwt),
		Expect().Status().Equal(http.StatusNoContent),
	)
}
