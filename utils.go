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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	BuyExperience  OnRampExperience = "buy"
	SendExperience OnRampExperience = "send"
	DefaultHost    string           = "https://pay.coinbase.com"
)

func SetCredentials() (*Credentials, error) {
	appId, ok := os.LookupEnv("CBPAY_APP_ID")
	if !ok {
		return nil, errors.New("environment variable CBPAY-APP-ID not set")
	}
	apiKey, ok := os.LookupEnv("CBPAY_API_KEY")
	if !ok {
		return nil, errors.New("environment variable CBPAY-API-KEY not set")

	}
	return &Credentials{
		AppId:  appId,
		ApiKey: apiKey,
	}, nil
}

func (c *Client) SetOptionsParams(r *http.Request, countryCode string) *http.Request {
	q := r.URL.Query()
	q.Add("country", countryCode)
	r.URL.RawQuery = q.Encode()

	return r
}

func (c *Client) SetOptionsWithSubdivision(r *http.Request, countryCode string, subdivision string) *http.Request {
	q := r.URL.Query()
	q.Add("country", countryCode)
	q.Add("subdivision", subdivision)
	r.URL.RawQuery = q.Encode()

	return r
}

func (c *Client) BuildTransactionUrl(params *TransactionRequest) string {
	baseUrl := fmt.Sprintf(c.HttpBaseUrl)
	userPart := fmt.Sprintf("/user/%s/transactions", url.PathEscape(params.PartnerUserId))
	v := url.Values{}

	v.Set("page_key", GetPageKey(params.PageKey, "1"))
	v.Set("page_size", GetPageSize(params.PageSize))

	return baseUrl + userPart + "?" + v.Encode()
}

func (c *Client) ValidateQuoteParams(params *BuyQuotePayload) error {

	if params == nil {
		return errors.New("BuyQuotePayload cannot be nil")
	}

	if params.PurchaseCurrency == "" {
		return errors.New("PurchaseCurrency cannot be empty")
	}

	if params.PaymentAmount == "" {

		return errors.New("PaymentAmount cannot be empty")
	}

	if params.PaymentCurrency == "" {
		return errors.New("PaymentCurrency cannot be empty")
	}

	if params.PaymentMethod == "" {
		return errors.New("PaymentMethod cannot be empty")
	}

	if params.Country == "" {

		return errors.New("country cannot be empty")
	}

	return nil
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("API error: Code: %d Message %s", e.Code, e.Message)
}

func handleApiResponse(resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	apiError := &ApiError{}
	apiError.Code = resp.StatusCode

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		apiError.Message = "failed to read response body"
		return apiError
	}

	if err := json.Unmarshal(body, apiError); err != nil {
		apiError.Message = "failed to parse response"
		return apiError
	}

	if apiError.Message == "" {
		apiError.Message = "Unknown API error occurred"
	}

	return apiError

}

func GetPageSize(pageSize *int) string {
	if pageSize == nil || *pageSize < 0 {
		return "1"
	}
	return strconv.Itoa(*pageSize)
}

func GetPageKey(ptr *string, defaultValue string) string {
	if ptr != nil {
		return *ptr
	}
	return defaultValue
}

func (c *Client) post(ctx context.Context, url string, p io.Reader) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, p)
	if err != nil {
		return nil, err
	}
	c.SetHeaders(req)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) get(ctx context.Context, url string, b io.Reader) ([]byte, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, b)
	if err != nil {
		return nil, err
	}

	c.SetHeaders(req)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
