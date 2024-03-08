package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func NewProduct(id string, name string, desc string, price float64) *Product {
	return &Product{id, name, desc, price}
}

func main() {
	createdProduct := getProduct()

	createdProduct.printData()
}

func getProduct() *Product {
	fmt.Println("Please enter the product data:")
	fmt.Println("------------------------------")

	reader := bufio.NewReader(os.Stdin) // returns pointer

	inputId := readUserInput(reader, "Product ID: ")
	inputTitle := readUserInput(reader, "Product Title: ")
	inputDescription := readUserInput(reader, "Product Description: ")
	inputPrice := readUserInput(reader, "Product Price: ")

	floatedInputPrice, _ := strconv.ParseFloat(inputPrice, 64)

	product := NewProduct(inputId, inputTitle, inputDescription, floatedInputPrice)
	return product
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}
