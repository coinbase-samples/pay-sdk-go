package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/juantmoore/pay-sdk"
)

func TestBuyConfig(t *testing.T) {

	//Arrange
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}
	c := pay.NewClient(creds, http.Client{})
	buyConfig := pay.ConfigData{}

	//Act
	config, err := c.BuyConfig(ctx)

	//Assert
	if err != nil {
		t.Fatalf("BuyConfig returned an unexpected error: %v", err)
	}

	if err := json.Unmarshal(config, &buyConfig); err != nil {
		t.Fatalf("error unmarshalling: %s ", err)
	}

	if len(buyConfig.Countries) == 0 {
		t.Errorf("Expected BuyConfig to return at least one country, got none")
	}

}
