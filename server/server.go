package main

import (
	"net/http"

	"github.com/Guilherme-daCosta/client-server-api/server/controller"
	"github.com/Guilherme-daCosta/client-server-api/server/db"
	"github.com/Guilherme-daCosta/client-server-api/server/model"
	"github.com/Guilherme-daCosta/client-server-api/server/repository"
	"github.com/Guilherme-daCosta/client-server-api/server/service"
)

func main() {
	db, err := db.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.AutoMigrate(&model.CurrencyQuote{})
	if err != nil {
		panic(err)
	}

	repo := repository.NewCurrencyQuoteRepository(db.DB)
	service := service.NewCurrencyQuoteService(repo)
	controller := controller.NewCurencyQuoteController(service)

	server := http.NewServeMux()
	server.HandleFunc("/cotacao", controller.GetCurrencyQuote)
	err = http.ListenAndServe(":8080", server)
	if err != nil {
		panic(err)
	}
}
