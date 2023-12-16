package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Guilherme-daCosta/client-server-api/server/dto"
	"github.com/Guilherme-daCosta/client-server-api/server/model"
)

type CurrencyQuoteService struct {
	repo model.CurrencyQuoteRepository
}

func NewCurrencyQuoteService(repo model.CurrencyQuoteRepository) *CurrencyQuoteService {
	return &CurrencyQuoteService{
		repo: repo,
	}
}

func (svc *CurrencyQuoteService) GetCurrencyQuoteAndSave() (*dto.CurrencyQuoteResponseDTO, error) {
	currencyQuoteDTO, err := svc.getCurrencyQuoteFromAPI()
	if err != nil {
		return nil, err
	}

	currencyQuote := &model.CurrencyQuote{
		Code: currencyQuoteDTO.USDBRL.Code,
		Codein: currencyQuoteDTO.USDBRL.Codein,
		Name: currencyQuoteDTO.USDBRL.Name,
		High: currencyQuoteDTO.USDBRL.High,
		Low: currencyQuoteDTO.USDBRL.Low,
		VarBid: currencyQuoteDTO.USDBRL.VarBid,
		PctChange: currencyQuoteDTO.USDBRL.PctChange,
		Bid: currencyQuoteDTO.USDBRL.Bid,
		Ask: currencyQuoteDTO.USDBRL.Ask,
		Timestamp: currencyQuoteDTO.USDBRL.Timestamp,
		CreateDate: currencyQuoteDTO.USDBRL.CreateDate,
	}

	err = svc.save(currencyQuote)
	if err != nil {
		return nil, err
	}

	resp := &dto.CurrencyQuoteResponseDTO{
		Bid: currencyQuote.Bid,
	}

	return resp, nil
}

func (svc *CurrencyQuoteService) getCurrencyQuoteFromAPI() (*dto.CurrencyQuoteDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var c dto.CurrencyQuoteDTO
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (svc *CurrencyQuoteService) save(currencyQuote *model.CurrencyQuote) error {
	return svc.repo.Save(currencyQuote)
}
