package main

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/pay-sdk-go"
)

func TestBuyOptions(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}
	c := pay.NewClient(creds, http.Client{})

	cc := "US"
	subdivision := "NY"

	response, err := c.BuyOptions(ctx, cc, &subdivision)
	if err != nil {
		t.Fatalf("error retrieving Buy Options")
	}

	if response.PaymentCurrencies == nil && response.PurchaseCurrencies == nil {
		t.Fatalf("error buy response returned nil")
	}
}
