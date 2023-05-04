package integration_test

import (
	"fmt"
	. "github.com/Eun/go-hit"
	base "github.com/alexgaas/order-reward/integration-test"
	orders "github.com/alexgaas/order-reward/internal/usecase/orders"
	"github.com/google/uuid"
	"net/http"
	"testing"
)

func TestGetOrders_OrdersNotFound(t *testing.T) {
	jwt := RegisterUser()

	// test orders not found
	Test(t,
		Description("Retrieve a list of user order numbers, including their processing status and accrual information"),
		Get(base.BasePath+"orders"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Headers("Authorization").Add(jwt),
		Expect().Status().Equal(http.StatusNoContent),
	)
}

func TestPostOrderHappyPath(t *testing.T) {
	jwt := RegisterUser()

	validNumber := orders.Generate(11)

	Test(t,
		Description("Add an order number for accrual operations"),
		Post(base.BasePath+"orders"),
		Send().Headers("Content-Type").Add("text/plain"),
		Send().Headers("Authorization").Add(jwt),
		Send().Body().String(validNumber),
		Expect().Status().Equal(http.StatusAccepted),
	)
}

func TestPostOrder_InvalidOrderFormat(t *testing.T) {
	jwt := RegisterUser()

	Test(t,
		Description("Add an order number for accrual operations"),
		Post(base.BasePath+"orders"),
		Send().Headers("Content-Type").Add("text/plain"),
		Send().Headers("Authorization").Add(jwt),
		Send().Body().String("987654321"),
		Expect().Status().Equal(http.StatusInternalServerError),
		Expect().Body().String().Contains("order number is not valid"),
	)
}

func RegisterUser() string {
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
	return jwt
}
