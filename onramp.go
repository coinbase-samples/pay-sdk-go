package pay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// Returns a URL that can be launched by a browser (a link or button on web, or a webview, or linking on mobile)
func (c *Client) GenerateOnRampUrl(ctx context.Context, p GenerateOnRampUrlOptions) (string, error) {

	if c.Credentials.AppId == "" {
		return "", errors.New("AppId not set")
	}

	host := DefaultHost
	if p.Host != nil {
		host = *p.Host
	}

	parsedUrl, err := url.Parse(host)
	if err != nil {
		return "", fmt.Errorf("%s \n host: %s", err, host)
	}
	parsedUrl.Path = "/buy/select-asset"

	destinationWallets, err := json.Marshal(p.OnRampAppParams.DestinationWallets)
	if err != nil {
		return "", err
	}

	v := url.Values{}
	v.Set("appId", c.Credentials.AppId)
	v.Set("destinationWallets", string(destinationWallets))

	parsedUrl.RawQuery = v.Encode()

	return parsedUrl.String(), nil
}
