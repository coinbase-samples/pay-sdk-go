package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/coinbase-samples/pay-sdk-go"
)

func TestSetHeaders(t *testing.T) {

	//Arrange
	appId := "123"
	apiKey := "456"
	url := ""
	creds := &pay.Credentials{
		ApiKey: apiKey,
		AppId:  appId,
	}

	req, err := http.NewRequest(url, http.MethodGet, nil)
	if err != nil {
		fmt.Print(err)
	}

	c := pay.NewClient(creds, http.Client{})
	fmt.Printf("apiKey before: %s \n", c.Credentials.ApiKey)
	fmt.Printf("appId before: %s \n", c.Credentials.AppId)

	//Act
	c.SetHeaders(req)
	fmt.Printf("apiKey: %s \n", req.Header.Get("CBPAY-API-KEY"))
	fmt.Printf("appId: %s \n", req.Header.Get("CBPAY-APP-ID"))

	//Assert

	if req.Header.Get("CBPAY-APP-ID") != appId {
		t.Errorf("Expected CBPAY_API_KEY to be %s, got %s", appId, req.Header.Get("CBPAY-APP-ID"))
	}

	if req.Header.Get("CBPAY-API-KEY") != apiKey {
		t.Errorf("Expected CBPAY_API_KEY to be %s, got %s", apiKey, req.Header.Get("CBPAY-API-KEY"))
	}
}

func TestBaseUrl(t *testing.T) {

	//Arrange
	creds := &pay.Credentials{
		ApiKey: "123",
		AppId:  "456",
	}

	c := pay.NewClient(creds, http.Client{})
	url := "abc"
	c.BaseUrl(url)
	fmt.Print(c.HttpBaseUrl)

	if c.HttpBaseUrl != url {
		t.Errorf("Expected BaseUrl to be %s got %s instead", url, c.HttpBaseUrl)
	}
}
