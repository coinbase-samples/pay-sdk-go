package main

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/pay-sdk-go"
)

func TestBuyQuote(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}
	c := pay.NewClient(creds, http.Client{})
	network := "Bitcoin"
	subdivision := "NY"
	payload := &pay.BuyQuotePayload{
		PurchaseCurrency: "BTC",
		PurchaseNetwork:  &network,
		PaymentAmount:    "100",
		PaymentCurrency:  "USD",
		PaymentMethod:    "CARD",
		Country:          "US",
		Subdivision:      &subdivision,
	}

	response, err := c.BuyQuote(ctx, payload)
	if err != nil {
		t.Fatalf("error retrieving Buy Quote: %s", err)
	}

	if response.QuoteId == "" {
		t.Fatalf("buy quote returned no quoteId")
	}

}
