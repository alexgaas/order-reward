package integration_test

import (
	"encoding/json"
	"fmt"
	base "github.com/alexgaas/order-reward/integration-test"
	"github.com/alexgaas/order-reward/internal/domain"
	"github.com/stretchr/testify/require"
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

func TestRegister_User_Already_Exists(t *testing.T) {
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

	Test(t,
		Description("User registration"),
		Post(base.BasePath+"user/register"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusConflict),
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
		Description("User registration"),
		Post(base.BasePath+"user/register"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
	)

	var res string

	MustDo(
		Description("User login"),
		Post(base.BasePath+"user/login"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
		Store().Response().Body().String().In(&res),
	)

	var loginResponse domain.LoginResponse
	_ = json.Unmarshal([]byte(res), &loginResponse)

	require.NotEmpty(t, loginResponse.Authtoken)
}

func TestLogin_Invalid_User(t *testing.T) {
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

	incorrectPassword := "incorrect"
	loginBody := fmt.Sprintf(`{
		"login": "test_%s",
		"password": "test_%s"
	}`, randomLogin, incorrectPassword)

	Test(t,
		Description("User login"),
		Post(base.BasePath+"user/login"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(loginBody),
		Expect().Status().Equal(http.StatusUnauthorized),
	)
}
