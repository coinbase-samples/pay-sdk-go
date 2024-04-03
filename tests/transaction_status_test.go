package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/pay-sdk-go"
)

func TestTransactionStatus(t *testing.T) {

	puid := "145"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}
	c := pay.NewClient(creds, http.Client{})

	tx := &pay.TransactionRequest{
		PartnerUserId: puid,
	}

	resp, err := c.TransactionStatus(ctx, tx)
	if err != nil {
		t.Fatalf("error receiving transaction status: %s", err)
	}

	fmt.Printf("transaction status: %s", resp)
}
