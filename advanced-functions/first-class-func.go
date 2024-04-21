package main

import "fmt"

type transformFn func(int) int // defining custom type for function type to pass in param
type anotherFn func(int, []string, map[string][]int) ([]int, string)

func firstclass() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := doubleNumbers(&numbers)            // REMINDER: we use pointers to avoid brand new copies of slice every time we use it
	tripled := transformNumbers(&numbers, triple) // IMPORTANT: since functions are VALUES in golang we can pass it through another function as parameter
	fmt.Println(doubled, tripled)

	moreNumbers := []int{5, 1, 7}
	transformerFn1 := transformerFunction(&numbers)
	transformerFn2 := transformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers, moreTransformedNumbers)
}

func doubleNumbers(numbers *[]int) []int {
	doubledNumbers := []int{}

	for _, value := range *numbers {
		// doubledNumbers = append(doubledNumbers, value*2)
		doubledNumbers = append(doubledNumbers, double(value))
	}

	return doubledNumbers
}

// ------ we can pass function in another function params:
func transformNumbers(numbers *[]int, transform transformFn) []int {
	transformedNumbers := []int{}

	for _, value := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(value))
	}

	return transformedNumbers
}

// ------ we can also return function in another function(used in closures):
func transformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double // we should not execute function (double()) and just pass the function through it
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
