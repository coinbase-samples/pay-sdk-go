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
	"fmt"
	"net/http"
	"testing"

	"github.com/coinbase-samples/pay-sdk-go"
)

func TestSetHeaders(t *testing.T) {

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

	c.SetHeaders(req)

	if req.Header.Get("CBPAY-APP-ID") != appId {
		t.Errorf("Expected CBPAY_API_KEY to be %s, got %s", appId, req.Header.Get("CBPAY-APP-ID"))
	}

	if req.Header.Get("CBPAY-API-KEY") != apiKey {
		t.Errorf("Expected CBPAY_API_KEY to be %s, got %s", apiKey, req.Header.Get("CBPAY-API-KEY"))
	}
}

func TestBaseUrl(t *testing.T) {

	creds := &pay.Credentials{
		ApiKey: "123",
		AppId:  "456",
	}

	c := pay.NewClient(creds, http.Client{})
	url := "abc"
	c.BaseUrl(url)

	if c.HttpBaseUrl != url {
		t.Errorf("Expected BaseUrl to be %s got %s instead", url, c.HttpBaseUrl)
	}
}
