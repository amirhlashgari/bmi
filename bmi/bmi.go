package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amirhlashgari/bmi/variable"
)

func main() {
	fmt.Println(variable.MainTitle)
	fmt.Println(variable.Separator)

	fmt.Print("Please enter your weight(kg): ")
	userWeight, _ := reader.ReadString('\n')

	fmt.Print("Please enter your height(m): ")
	userHeight, _ := reader.ReadString('\n')

	userWeight = strings.Replace(userWeight, "\n", "", -1)
	userHeight = strings.Replace(userHeight, "\n", "", -1)

	weight, _ := strconv.ParseFloat(userWeight, 64)
	height, _ := strconv.ParseFloat(userHeight, 64)

	bmi := weight / (height * height)

	fmt.Printf("Your BMI: %.2f", bmi) // %.2f specifies that show only two decimal of float64
	fmt.Println(variable.EmptyString)
	fmt.Println(variable.Separator)
}
