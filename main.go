package main

import (
	"fmt"
	"strings"
)

const (
	pricePerItem = 800000 // Price of each item in IRR
)

func calculateFinalPrice(items []string) int {
	itemCount := make(map[string]int)
	for _, item := range items {
		itemCount[item]++
	}

	// Create a slice to hold the counts of unique items
	var counts []int
	for _, count := range itemCount {
		counts = append(counts, count)
	}

	// This will hold the maximum possible price
	var maxPrice int

	// A function to calculate the price with a specific combination of unique item types
	var calculatePrice func([]int, int, float64)
	calculatePrice = func(counts []int, totalItems int, currentDiscount float64) {
		if totalItems == 0 {
			// Calculate final price after discounts
			price := totalItems * pricePerItem
			price = int(float64(price) * (1 - currentDiscount))
			if price > maxPrice {
				maxPrice = price
			}
			return
		}

		for i := 0; i < len(counts); i++ {
			if counts[i] > 0 {
				// Create a copy of counts to avoid mutating the original
				newCounts := make([]int, len(counts))
				copy(newCounts, counts)

				// Decrease the count of the selected type
				newCounts[i]--

				// Calculate the number of unique types we are using
				uniqueTypes := 0
				for _, count := range newCounts {
					if count > 0 {
						uniqueTypes++
					}
				}

				// Apply discount based on the number of unique types
				var discount float64
				switch uniqueTypes {
				case 5:
					discount = 0.25
				case 4:
					discount = 0.20
				case 3:
					discount = 0.10
				case 2:
					discount = 0.05
				}

				// Recurse with the new counts and updated discount
				calculatePrice(newCounts, totalItems-1, currentDiscount+discount)
			}
		}
	}

	// Start the recursion
	calculatePrice(counts, len(items), 0)

	return maxPrice
}

func main() {
	input := "shirt,shirt,pants,pants,jacket,jacket,shoes,hoody"
	items := strings.Split(input, ",")
	finalPrice := calculateFinalPrice(items)
	fmt.Println(finalPrice) // Should output 51200000
}
