package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type CurrencyQuote struct {
	Bid string `json:"bid"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	endpoint := "http://localhost:8080/cotacao"
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == http.StatusOK {
		var quotation CurrencyQuote
		err = json.Unmarshal(body, &quotation)
		if err != nil {
			panic(err)
		}

		err = saveCurrencyQuote(&quotation)
		if err != nil {
			panic(err)
		}

		log.Println("quotation saved successful")
		return
	}

	var errResp ErrorResponse
	err = json.Unmarshal(body, &errResp)
	if err != nil {
		panic(err)
	}

	log.Printf("error: %s", errResp.Message)
}

func saveCurrencyQuote(currencyQuote *CurrencyQuote) error {
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	text := fmt.Sprintf("Dólar: %s\n", currencyQuote.Bid)
	_, err = file.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}
