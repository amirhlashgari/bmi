package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// instead of defining double and triple function, we can pass it in params as an anonymos function.
	// the anonymos function params and return should sync with what defined in transformNumbers2
	transformed := transformNumbers2(&numbers, func(number int) int { return number * 2 })

	fmt.Println(transformed)
}

func transformNumbers2(numbers *[]int, transform func(int) int) []int {
	transformedNumbers := []int{}

	for _, value := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(value))
	}

	return transformedNumbers
}
