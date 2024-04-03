package main

import (
	"context"
	"encoding/json"
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
