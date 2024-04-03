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
