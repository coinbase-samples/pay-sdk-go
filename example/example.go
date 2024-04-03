package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/coinbase-samples/pay-sdk-go"
)

func gogo() {

	creds, err := pay.SetCredentials()
	if err != nil {
		fmt.Print(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := pay.NewClient(creds, http.Client{})

	resp, err := client.BuyConfig(ctx)
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	config := &pay.BuyConfigResponse{}
	if err = json.Unmarshal(resp, config); err != nil {
		fmt.Printf("error: %s", err)
	}

	fmt.Printf("response: %v", config)
}
