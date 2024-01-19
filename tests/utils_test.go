package main

//fix error messages
//fix variable names like resp, v, cc, countryCode, confirmUrl

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/pay-sdk-go"
)

func TestSetOptionsParams(t *testing.T) {
	//Arrange
	confirmUrl := "https://pay.coinbase.com/api/v1/buy/options?country=US"
	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}

	c := pay.NewClient(creds, http.Client{})
	url := c.HttpBaseUrl + "/options"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	country := "US"

	//Act
	c.SetOptionsParams(req, country)

	//Assert
	if req.URL.String() != confirmUrl {
		t.Fatalf("unexpected url: got %s, expected %s", req.URL.String(), confirmUrl)
	}

}

func TestSetOptionsWithSubdivision(t *testing.T) {

	//Arrange
	countryCode := "US"
	subdivision := "NY"
	confirmUrl := fmt.Sprintf("https://pay.coinbase.com/api/v1/buy/options?country=%s&subdivision=%s", countryCode, subdivision)
	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}

	c := pay.NewClient(creds, http.Client{})
	url := c.HttpBaseUrl + "/options"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	country := "US"
	sub := "NY"

	//Act
	c.SetOptionsWithSubdivision(req, country, sub)

	//Assert
	if req.URL.String() != confirmUrl {
		t.Fatalf("unexpected url: got %s, expected %s", req.URL.String(), confirmUrl)
	}
}

func TestBuildTransactionUrl(t *testing.T) {

	userId := "1234-5678"
	expectedUrl := "https://pay.coinbase.com/api/v1/buy/user/1234-5678/transactions?page_key=1&page_size=1"
	pageKey := "1"
	pageSize := int(1)

	params := &pay.TransactionRequest{
		PartnerUserId: userId,
		PageKey:       &pageKey,
		PageSize:      &pageSize,
	}

	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}

	c := pay.NewClient(creds, http.Client{})
	result := c.BuildTransactionUrl(params)
	expectedUrlParsed, _ := url.Parse(expectedUrl)
	resultParsed, _ := url.Parse(result)

	if resultParsed.Path != expectedUrlParsed.Path {
		t.Fatalf("unexpected path: got %s, expected %s", resultParsed.Path, expectedUrlParsed.Path)
	}

	if resultParsed.RawQuery != expectedUrlParsed.RawQuery {
		t.Fatalf("unexpected query: got %s, expected %s", resultParsed.RawQuery, expectedUrlParsed.RawQuery)

	}

}
