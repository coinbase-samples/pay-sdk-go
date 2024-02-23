package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/pay-sdk-go"
)

func TestGenerateOnRampUrl(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}
	c := pay.NewClient(creds, http.Client{})

	destinationWallets := []pay.DestinationWallet{{
		Address:           "0xabcdef",
		Blockchains:       &[]string{"solana", "ethereum"},
		Assets:            &[]string{"USDC", "ETH"},
		SupportedNetworks: &[]string{"USDC", "ethereum"}},
	}

	onRampParams := pay.OnRampAppParams{
		DestinationWallets: destinationWallets,
	}

	p := pay.GenerateOnRampUrlOptions{
		AppId:           c.Credentials.AppId,
		Host:            &c.Host,
		OnRampAppParams: onRampParams,
	}

	parsedUrl, _ := url.Parse(c.Host)
	parsedUrl.Path = "/buy/select-asset"
	destinationWalletsJson, _ := json.Marshal(destinationWallets)
	v := url.Values{}
	v.Set("appId", c.Credentials.AppId)
	v.Set("destinationWallets", string(destinationWalletsJson))
	expectedUrl := parsedUrl.String() + "?" + v.Encode()

	actualUrl, err := c.GenerateOnRampUrl(ctx, p)
	if err != nil {
		t.Fatalf("err %s", err)
	}

	if err != nil {
		t.Errorf("GenerateOnRampUrl returned an error: %v", err)
	}
	if actualUrl != expectedUrl {
		t.Errorf("GenerateOnRampUrl returned unexpected URL: got %v, want %v", actualUrl, expectedUrl)
	}

}
