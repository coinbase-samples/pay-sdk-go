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
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/coinbase-samples/pay-sdk-go"
)

func TestBuyQuote(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}
	c := pay.NewClient(creds, http.Client{})
	network := "Bitcoin"
	subdivision := "NY"
	payload := &pay.BuyQuotePayload{
		PurchaseCurrency: "BTC",
		PurchaseNetwork:  &network,
		PaymentAmount:    "100",
		PaymentCurrency:  "USD",
		PaymentMethod:    "CARD",
		Country:          "US",
		Subdivision:      &subdivision,
	}

	response, err := c.BuyQuote(ctx, payload)
	if err != nil {
		t.Fatalf("error retrieving Buy Quote: %s", err)
	}

	if response.QuoteId == "" {
		t.Fatalf("buy quote returned no quoteId")
	}

}
