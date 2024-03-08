package main

import "fmt"

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func (prod *Product) printData() {
	fmt.Printf("ID: %v\n", (*prod).id) // no need to destruct Product to use its attributes, (*prod).id === prod.id
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price: USD %.2f\n\n", prod.price)
}

func newProduct(id string, name string, desc string, price float64) *Product {
	return &Product{id, name, desc, price}
}

func main() {
	firstProduct := Product{"123456789", "A Book", "book's description", 6.99} // simple type of defining new product
	secondProduct := newProduct("987654321", "A Carpet", "nice one", 65.90)    // define product using struct methods

	firstProduct.printData()
	secondProduct.printData()
}
