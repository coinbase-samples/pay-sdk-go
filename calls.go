package pay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Returns the list of countries supported by Coinbase Pay, and the payment methods available in each country
func (c *Client) BuyConfig(ctx context.Context) ([]byte, error) {

	url := fmt.Sprintf(c.HttpBaseUrl + "/config")

	body, err := c.get(ctx, url, nil)
	if err != nil {
		return nil, err
	}

	return body, nil

}

// Returns the supported fiat currencies and available crypto assets that can be passed into the Buy Quote API
func (c *Client) BuyOptions(ctx context.Context, countryCode string, subdivision *string) (*BuyOptionsResponse, error) {

	url := fmt.Sprintf(c.HttpBaseUrl + "/options")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if subdivision != nil {
		req = c.SetOptionsWithSubdivision(req, countryCode, *subdivision)
	} else {
		req = c.SetOptionsParams(req, countryCode)
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

	apiResponse := &BuyOptionsResponse{}
	if err = json.Unmarshal(body, apiResponse); err != nil {
		return nil, err
	}

	return apiResponse, nil
}

// Provides a quote based on the asset the user would like to purchase and other parameters
func (c *Client) BuyQuote(ctx context.Context, quoteParams *BuyQuotePayload) (*BuyQuoteResponse, error) {

	if err := c.ValidateQuoteParams(quoteParams); err != nil {
		return nil, err
	}

	url := fmt.Sprintf(c.HttpBaseUrl + "/quote")
	payload, err := json.Marshal(quoteParams)
	if err != nil {
		return nil, err
	}

	body, err := c.post(ctx, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	apiResponse := &BuyQuoteResponse{}
	if err = json.Unmarshal(body, apiResponse); err != nil {
		return nil, err
	}

	return apiResponse, nil

}

// Provides clients with a list of the user's CBPay transactions
func (c *Client) TransactionStatus(ctx context.Context, t *TransactionRequest) (*TransactionResponse, error) {

	url := c.BuildTransactionUrl(t)
	body, err := c.get(ctx, url, nil)
	if err != nil {
		return nil, err
	}

	apiResponse := &TransactionResponse{}
	if err = json.Unmarshal(body, apiResponse); err != nil {
		return nil, err
	}

	return apiResponse, nil
}

// Returns a session token as a secure way for the client to initialize the Pay SDK
func (c *Client) GetSessionToken(ctx context.Context, d *DestinationWallet) (*Token, error) {

	url := "https://pay.coinbase.com/api/v1/onramp/token"
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(d)
	if err != nil {
		return nil, err
	}

	body, err := c.post(ctx, url, buf)
	if err != nil {
		return nil, err
	}

	token := &Token{}
	if err := json.Unmarshal(body, token); err != nil {
		return nil, err
	}

	return token, err

}
