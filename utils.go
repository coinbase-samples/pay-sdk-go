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
	//Should be optional?
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

	if params.PageKey != nil {
		v.Set("page_key", *params.PageKey)
	}

	v.Set("page_size", GetPageSize(params.PageSize))

	return baseUrl + userPart + "?" + v.Encode()
}

func (c *Client) GenerateOnRampUrl(ctx context.Context, p GenerateOnRampUrlOptions) (string, error) {

	if c.Credentials.AppId == "" {
		return "", errors.New("error: AppId not set")
	}

	host := DefaultHost
	if p.Host != nil {
		host = *p.Host
	}

	parsedUrl, err := url.Parse(host)
	if err != nil {
		return "", err
	}
	parsedUrl.Path = "/buy/select-asset"

	destinationWallets, err := json.Marshal(p.OnRampAppParams.DestinationWallets)
	if err != nil {
		return "", err
	}

	v := url.Values{}
	v.Set("appId", c.Credentials.AppId)
	v.Set("destinationWallets", string(destinationWallets))

	//sometimes p.OnrampAppParams will have some options fields incluced.
	//How do I check for this to include them in the parameters when generating the url?

	parsedUrl.RawQuery = v.Encode()
	fmt.Println(parsedUrl.String())
	return parsedUrl.String(), nil
}

func (c *Client) ValidateQuoteParams(params *BuyQuotePayload) error {
	//add validation for optional fields?

	if params == nil {
		return errors.New("BuyQuotePayload cannot be nil")
	}

	if params.PurchaseCurrency == "" {
		return errors.New("PurchaseCurrency cannot be empty")
	}

	if params.PaymentAmount == "" {
		//validate format?
		return errors.New("PaymentAmount cannot be empty")
	}

	if params.PaymentCurrency == "" {
		return errors.New("PaymentCurrency cannot be empty")
	}

	if params.PaymentMethod == "" {
		return errors.New("PaymentMethod cannot be empty")
	}

	if params.Country == "" {
		//validate ISO?
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
