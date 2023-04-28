package main

import (
	"github.com/alexgaas/order-reward/internal/api"
	"github.com/alexgaas/order-reward/internal/config"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "ORDER REWARD:\t", log.Ldate|log.Ltime)

	appConfig, err := config.GetNewConfig(config.GetAppFlags())
	if err != nil {
		logger.Fatalln(err)
	}

	app, err := api.NewAppHandler(*appConfig, logger)
	if err != nil {
		logger.Fatalln(err)
	}
	defer app.Storage.CloseDB()

	if err := app.Storage.InitDB(); err != nil {
		logger.Fatalln(err)
	}

	router := api.NewRouter(app)
	logger.Println("App is waiting connections on: ", appConfig.AppAddress)
	logger.Fatal(http.ListenAndServe(appConfig.AppAddress, router))
}
