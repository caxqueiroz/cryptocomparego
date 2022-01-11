package main

import (
	"fmt"

	"context"
	"github.com/caxqueiroz/cryptocomparego"
)

func main() {

	client := cryptocomparego.NewClient(nil)
	ctx := context.TODO()

	coinList, _, err := client.Coin.List(ctx)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
	}

	for _, coin := range coinList {
		fmt.Printf("Coin %s - %s\n", coin.Name, coin.FullName)
	}
}
