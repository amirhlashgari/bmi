package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("BMI Calculator")
	fmt.Println("--------------------")

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
}
