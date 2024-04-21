package main

import "fmt"

func main() {
	productPrices := map[string]float64{
		"Carpet": 99.50,
		"Book":   21.99,
		"Ball":   10.10,
	}

	fmt.Println(productPrices)
	fmt.Println(productPrices["Book"])

	productPrices["Chair"] = 23.99
	fmt.Println(productPrices)

	delete(productPrices, "Ball")
	fmt.Println(productPrices)

}
