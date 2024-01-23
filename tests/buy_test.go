package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/pay-sdk-go"
)

func TestBuyConfig(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}
	c := pay.NewClient(creds, http.Client{})
	buyConfig := pay.ConfigData{}

	config, err := c.BuyConfig(ctx)

	if err != nil {
		t.Fatalf("BuyConfig returned an unexpected error: %v", err)
	}

	if err := json.Unmarshal(config, &buyConfig); err != nil {
		t.Fatalf("error unmarshalling: %s ", err)
	}

	if len(buyConfig.Countries) == 0 {
		t.Fatal("error expected BuyConfig to return at least one country, got none")
	}

}

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

func TestTransactioNStatus(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}
	c := pay.NewClient(creds, http.Client{})

	tx := &pay.TransactionRequest{
		PartnerUserId: "",
	}

	resp, err := c.TransactionStatus(ctx, tx)
	if err != nil {
		t.Fatalf("error receiving transaction status: %s", err)
	}

	fmt.Printf("response: %s", resp)
}
