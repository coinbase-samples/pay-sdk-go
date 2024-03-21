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

package pay

import (
	"net/http"
)

var baseUrl = "https://pay.coinbase.com/api/v1/buy"

type Client struct {
	HttpClient  http.Client
	Credentials *Credentials
	HttpBaseUrl string
	Host        string
}

func NewClient(creds *Credentials, httpClient http.Client) *Client {

	return &Client{
		HttpClient:  httpClient,
		Credentials: creds,
		HttpBaseUrl: baseUrl,
		Host:        DefaultHost,
	}
}

func (c *Client) SetHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("CBPAY-VERSION", "2018-03-22")
	req.Header.Set("CBPAY-APP-ID", c.Credentials.AppId)
	req.Header.Set("CBPAY-API-KEY", c.Credentials.ApiKey)
}

func (c *Client) BaseUrl(u string) *Client {
	c.HttpBaseUrl = u
	return c

}
