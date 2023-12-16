package controller

import (
	"net/http"

	"github.com/Guilherme-daCosta/client-server-api/server/rest"
	"github.com/Guilherme-daCosta/client-server-api/server/service"
)

type currencyQuoteController struct {
	service *service.CurrencyQuoteService
}

func NewCurencyQuoteController(service *service.CurrencyQuoteService) *currencyQuoteController {
	return &currencyQuoteController{
		service: service,
	}
}

func (ctrl *currencyQuoteController) GetCurrencyQuote(w http.ResponseWriter, r *http.Request) {
	currencyQuote, err := ctrl.service.GetCurrencyQuoteAndSave()
	if err != nil {
		rest.NewBadRequest(w, err)
		return
	}

	rest.NewSucessful(w, currencyQuote)
}