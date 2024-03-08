package main

import "fmt"

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func newProduct(id string, name string, desc string, price float64) *Product {
	return &Product{id, name, desc, price}
}

func main() {
	firstProduct := Product{"123456789", "A Book", "book's description", 6.99} // simple type of defining new product
	secondProduct := newProduct("987654321", "A Carpet", "nice one", 65.90)    // define product using struct methods

	fmt.Println(firstProduct)
	fmt.Println(*secondProduct)
}
