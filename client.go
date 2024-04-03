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
