package main

import "fmt"

func variadic() {
	numbers := []int{1, 10, 13}
	sum := sumupNormal(numbers)
	sum2 := sumupVariadic(1, 10, 13, 15, 18, -5)
	fmt.Println(sum)
	fmt.Println(sum2)
}

func sumupNormal(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func sumupVariadic(startNum int, numbers ...int) int {
	sum := startNum // ---> here first number will go here and others passing in function call will go for numbers slice
	for _, val := range numbers {
		sum += val
	}
	return sum
}
