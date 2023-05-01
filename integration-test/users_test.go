package integration_test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	. "github.com/Eun/go-hit"
	"github.com/google/uuid"
)

const (
	host     = "localhost:8000"
	basePath = "http://" + host + "/api/user"
)

func TestMain(m *testing.M) {
	log.Printf("Integration tests: host %s is available", host)
	code := m.Run()
	os.Exit(code)
}

// handler - /api/user/register
func TestRegisterHappyPath(t *testing.T) {
	randomLogin := uuid.New()
	randomPassword := uuid.New()
	body := fmt.Sprintf(`{
		"login": "test_%s",
		"password": "test_%s"
	}`, randomLogin, randomPassword)
	Test(t,
		Description("User registration"),
		Post(basePath+"/register"),
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
		Post(basePath+"/login"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
	)
}
