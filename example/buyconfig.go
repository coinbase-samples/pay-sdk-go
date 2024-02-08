package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/coinbase-samples/pay-sdk-go"
)

func main() {
	// // Arrange
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// creds := &pay.Credentials{
	// 	ApiKey: os.Getenv("CBPAY_API_KEY"),
	// 	AppId:  os.Getenv("CBPAY_APP_ID"),
	// }
	// c := pay.NewClient(creds, http.Client{})
	// buyConfig := pay.ConfigData{}

	// // Act
	// config, err := c.BuyConfig(ctx)
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// configString := string(config)
	// fmt.Print(configString + "\n")

	// if err := json.Unmarshal(config, &buyConfig); err != nil {
	// 	fmt.Printf("error unmarshalling: %s ", err)
	// }
	// fmt.Printf("\ndata: %v", buyConfig)

	//Arrange
	creds := &pay.Credentials{
		ApiKey: os.Getenv("CBPAY_API_KEY"),
		AppId:  os.Getenv("CBPAY_APP_ID"),
	}

	c := pay.NewClient(creds, http.Client{})
	countryCode := "US"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Act
	//why did I choose to return an struct of type BuyOptions response
	v, err := c.BuyOptions(ctx, countryCode, nil)
	//Assert

	fmt.Print(v)

	if err != nil {
		log.Fatalf("err, %s", err)
	}

	if v.Data == nil {
		fmt.Printf("%#v", v)
		log.Fatalf("call failed: %s", err)
	}
	//wrong countryCode
}
