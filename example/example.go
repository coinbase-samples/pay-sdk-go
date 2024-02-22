package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/coinbase-samples/pay-sdk-go"
)

func main() {

	creds, err := pay.SetCredentials()
	if err != nil {
		fmt.Print(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//initiate client from the pay package
	c := pay.NewClient(creds, http.Client{})
	destinationAddress := "0x123"
	d := pay.DestinationWallet{
		Address:     destinationAddress,
		Blockchains: &[]string{"Ethereum", "Solana"},
		Assets:      &[]string{"USDC"},
	}
	p := pay.OnRampAppParams{
		DestinationWallets: []pay.DestinationWallet{d},
	}
	o := pay.GenerateOnRampUrlOptions{
		AppId:           c.Credentials.AppId,
		Host:            &c.Host,
		OnRampAppParams: p,
	}

	url, err := c.GenerateOnRampUrl(ctx, o)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Generated URL:", url)
}
