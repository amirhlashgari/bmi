package main

import "fmt"

func anonymous() {
	numbers := []int{1, 2, 3, 4, 5}

	// instead of defining double and triple function, we can pass it in params as an anonymos function.
	// the anonymos function params and return should sync with what defined in transformNumbers2
	transformed := transformNumbers2(&numbers, func(number int) int { return number * 2 })

	doubledClosure := createTransformerClosure(2)
	tripledClosure := createTransformerClosure(3)

	transformedDouble := transformNumbers2(&numbers, doubledClosure)
	transformedTriple := transformNumbers2(&numbers, tripledClosure)

	fmt.Println(transformed)
	fmt.Println(transformedDouble)
	fmt.Println(transformedTriple)
}

func transformNumbers2(numbers *[]int, transform func(int) int) []int {
	transformedNumbers := []int{}

	for _, value := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(value))
	}

	return transformedNumbers
}

// this function returns another function
func createTransformerClosure(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
