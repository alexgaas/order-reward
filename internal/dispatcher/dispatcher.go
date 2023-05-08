package dispatcher

import (
	"context"
	"fmt"
	"github.com/alexgaas/order-reward/internal/domain"
	repository "github.com/alexgaas/order-reward/internal/repo"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"time"
)

const (
	statNew        = "NEW"
	statProcessing = "PROCESSING"
	checkPause     = 5 * time.Second
)

type Dispatcher struct {
	Storage        *repository.Repository
	Logger         *log.Logger
	AccrualAddress string
}

func (dsp Dispatcher) Run(ctx context.Context) {
	answer := domain.AccrualAnswer{}
	checkNew := true
	client := resty.New()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		clientCtx, cltCancel := context.WithTimeout(ctx, 3*time.Second)
		defer cltCancel()
		databaseCtx, dbCancel := context.WithTimeout(ctx, time.Second)
		defer dbCancel()

		var status string
		if checkNew {
			status = statNew
		} else {
			status = statProcessing
		}

		orderNumbers, err := dsp.Storage.DispatchGetOrders(databaseCtx, status)
		if err != nil {
			dsp.Logger.Print(err)
		}

		if len(orderNumbers) == 0 {
			checkNew = !checkNew
			time.Sleep(checkPause)
			continue
		}

		dsp.Logger.Printf("Dispatcher, orderNums = %+v, type = %T", orderNumbers, orderNumbers)

		for _, order := range orderNumbers {

			resp, err := client.R().
				SetContext(clientCtx).
				SetBody(order).
				SetResult(&answer).
				Get(fmt.Sprint(dsp.AccrualAddress, "/orders/", order))

			if err != nil {
				dsp.Logger.Print(err)
				continue
			}

			if resp.StatusCode() != http.StatusOK {
				dsp.Logger.Printf("For order number = %v, accrual system returned status: %v", order, resp.StatusCode())
				dsp.Logger.Printf("Answer = '%+v'", resp.String())
				continue
			}

			switch answer.Status {
			case "INVALID", "PROCESSED":
				if err = dsp.Storage.DispatchUpdateOrder(
					databaseCtx,
					domain.Order{
						Number:  answer.OrderNumber,
						Status:  answer.Status,
						Accrual: answer.Accrual,
					},
				); err != nil {
					dsp.Logger.Print(err)
					continue
				}
			case "REGISTERED", "PROCESSING":
				continue
			}
		}

		checkNew = !checkNew
	}
}
