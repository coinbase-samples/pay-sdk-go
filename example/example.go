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
	"fmt"
	"net/http"
	"time"

	"github.com/coinbase-samples/pay-sdk-go"
)

func main() {

	creds, err := pay.SetCredentials()
	if err != nil {
		fmt.Print(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//initiate client from the pay package
	c := pay.NewClient(creds, http.Client{})
	destinationAddress := "0x123"
	d := pay.DestinationWallet{
		Address:     destinationAddress,
		Blockchains: &[]string{"Ethereum", "Solana"},
		Assets:      &[]string{"USDC"},
	}
	p := pay.OnRampAppParams{
		DestinationWallets: []pay.DestinationWallet{d},
	}
	o := pay.GenerateOnRampUrlOptions{
		AppId:           c.Credentials.AppId,
		Host:            &c.Host,
		OnRampAppParams: p,
	}

	url, err := c.GenerateOnRampUrl(ctx, o)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Generated URL:", url)
}
