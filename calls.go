package pay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) BuyConfig(ctx context.Context) ([]byte, error) {

	url := fmt.Sprintf(c.HttpBaseUrl + "/config")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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

// How would I make BuyOptions one function? Instead of two func, maybe pass a *BuyOptionsRequest
func (c *Client) BuyOptions(ctx context.Context, countryCode string) (*BuyOptionsResponse, error) {

	url := fmt.Sprintf(c.HttpBaseUrl + "/options")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = c.SetOptionsParams(req, countryCode)
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

func (c *Client) BuyOptionsWithSubdivision(ctx context.Context, countryCode string, subdivision string) (*BuyOptionsResponse, error) {

	url := fmt.Sprintf(c.HttpBaseUrl + "/buy/options")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = c.SetOptionsWithSubdivision(req, countryCode, subdivision)
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

func (c *Client) BuyQuote(ctx context.Context, quoteParams *BuyQuotePayload) (*BuyQuoteResponse, error) {

	if err := c.ValidateQuoteParams(quoteParams); err != nil {
		return nil, err
	}

	payload, err := json.Marshal(quoteParams)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(c.HttpBaseUrl + "/quote")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	c.SetHeaders(req)
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if err := handleApiResponse(resp); err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	apiResponse := &BuyQuoteResponse{}
	if err = json.Unmarshal(body, apiResponse); err != nil {
		return nil, err
	}

	return apiResponse, nil

}

func (c *Client) TransactionStatus(ctx context.Context, t *TransactionRequest) (*TransactionResponse, error) {

	url := c.BuildTransactionUrl(t)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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

	apiResponse := &TransactionResponse{}
	if err = json.Unmarshal(body, apiResponse); err != nil {
		return nil, err
	}

	return apiResponse, nil
}