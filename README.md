# Coinbase Pay Go SDK

## Overview

The Pay Go is an SDK to drive the Coinbase Pay [REST APIs](https://docs.cloud.coinbase.com/pay-sdk/docs/rest-api-overview).

## License

The Pay Go SDK sample library is free and open source and released under the Apache License, Version 2.0.

The application and code are only available for demonstration purposes.

## Usage

To use the Pay Go SDK, initialize the Credentials struct and create a new client. The Credentials struct is JSON enabled. Ensure that Pay API credentials are stored in a secure manner.

> [!IMPORTANT]  
> If you do not have an appId or API key this SDK will not work. To sign up for Coinbase Pay fill out the [interest form](https://www.coinbase.com/cloud/cloud-interest).

## Examples

### Initalize credentials

Be sure to set your environment variables with your Coinbase Pay appId and API key.

```bash
export CBPAY_APP_ID=<Your-appId>
export CBPAY_API_KEY=<Your-API-Key>
```

```go
func main() {

	creds, err := pay.SetCredentials()
	if err != nil {
		fmt.Print(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	c := pay.NewClient(creds, http.Client{})
}
```

### Generate an Pay onramp URL

```go
func main() {

	creds, err := pay.SetCredentials()
	if err != nil {
		fmt.Print(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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
```

### Get the list of countries supported by Coinbase Pay

```go
func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}
	c := pay.NewClient(creds, http.Client{})
	buyConfig := pay.ConfigData{}

	config, err := c.BuyConfig(ctx)
    if err != nil{
        fmt.Printf(err)
        return
    }

    fmt.Println(string(config))

	}
```
