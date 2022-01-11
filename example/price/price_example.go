package main

import (
	"context"
	"fmt"
	"github.com/caxqueiroz/cryptocomparego"
)

func main() {

	client := cryptocomparego.NewClient(nil)
	ctx := context.TODO()

	priceRequest := cryptocomparego.NewPriceRequest("ETH", []string{"BRL", "USD", "EUR"})
	priceList, _, err := client.Price.List(ctx, priceRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
	}

	for _, coin := range priceList {
		fmt.Printf("Coin %s - %f\n", coin.Name, coin.Value)
	}
}
