package main

import (
	"fmt"

	"mamba-mentality.com/api"
)

func main() {

	rate, err := api.GetRate("BTC");

	if(err == nil) {
		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price);
	};
};