package integration_test

import (
	"encoding/json"
	"fmt"
	"github.com/alexgaas/order-reward/internal/domain"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
	base "github.com/alexgaas/order-reward/integration-test"
)

// before starting that integration tests you have to:
/*
start hoverfly as webserver:
	hoverctl start webserver

set simulate mode:
	hoverctl mode simulate

import points API configuration:
	 hoverctl import hoverfly/points_api_config.json

validate with curl:
	curl http://localhost:8500/orders/12345678903
*/

func TestGetPointsAccrualHappyPath(t *testing.T) {
	var res string
	validOrderNumber := "12345678903"
	relativePath := fmt.Sprintf(`orders/%s`, validOrderNumber)

	MustDo(
		Description("Retrieves data about reward processing of accrual"),
		Get(base.AccrualPath+relativePath),
		Send().Headers("Content-Type").Add("application/json"),
		Expect().Status().Equal(http.StatusOK),
		Store().Response().Body().String().In(&res),
	)

	require.NotEmpty(t, res)

	answer := domain.AccrualAnswer{}
	_ = json.Unmarshal([]byte(res), &answer)

	require.Exactly(t, answer.OrderNumber, validOrderNumber)
}
