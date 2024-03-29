/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

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
	confirmUrl := "https://pay.coinbase.com/api/v1/buy/options?country=US"
	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}

	c := pay.NewClient(creds, http.Client{})
	url := c.HttpBaseUrl + "/buy/options"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	country := "US"

	c.SetOptionsParams(req, country)

	if req.URL.String() != confirmUrl {
		t.Fatalf("unexpected url: got %s, expected %s", req.URL.String(), confirmUrl)
	}

}

func TestSetOptionsWithSubdivision(t *testing.T) {

	countryCode := "US"
	subdivision := "NY"
	confirmUrl := fmt.Sprintf("https://pay.coinbase.com/api/v1/buy/options?country=%s&subdivision=%s", countryCode, subdivision)
	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}

	c := pay.NewClient(creds, http.Client{})
	url := c.HttpBaseUrl + "/buy/options"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	country := "US"
	sub := "NY"

	c.SetOptionsWithSubdivision(req, country, sub)

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
