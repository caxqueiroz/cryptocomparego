package main

import (
	"context"
	"fmt"
	"github.com/caxqueiroz/cryptocomparego"
	"log"
)

func main() {
	client := cryptocomparego.NewClient(nil)
	ctx := context.TODO()

	histominRequest := cryptocomparego.NewHistominuteRequest("BTC", "ETH", 20, 1559802780)
	data, _, err := client.Histomin.Get(ctx, histominRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
