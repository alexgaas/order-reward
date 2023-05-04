package integration_test

import (
	"fmt"
	base "github.com/alexgaas/order-reward/integration-test"
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
	"github.com/google/uuid"
)

func TestRegisterHappyPath(t *testing.T) {
	randomLogin := uuid.New()
	randomPassword := uuid.New()
	body := fmt.Sprintf(`{
		"login": "test_%s",
		"password": "test_%s"
	}`, randomLogin, randomPassword)
	Test(t,
		Description("User registration"),
		Post(base.BasePath+"user/register"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
	)
}

func TestLoginHappyPath(t *testing.T) {
	randomLogin := uuid.New()
	randomPassword := uuid.New()
	body := fmt.Sprintf(`{
		"login": "test_%s",
		"password": "test_%s"
	}`, randomLogin, randomPassword)
	Test(t,
		Description("User login"),
		Post(base.BasePath+"user/login"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
	)
}
