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

func TestBuyOptions(t *testing.T) {

	//Arrange
	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}

	c := pay.NewClient(creds, http.Client{})
	countryCode := "US"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Act
	//why did I choose to return an struct of type BuyOptions response
	v, err := c.BuyOptions(ctx, countryCode, nil)
	//Assert

	fmt.Print(v)

	if err != nil {
		t.Fatalf("err, %s", err)
	}

	if v.Data == nil {
		fmt.Printf("%#v", v)
		t.Fatalf("call failed: %s", err)
	}
	//wrong countryCode

}
