package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("BMI Calculator")
	fmt.Println("--------------------")

	fmt.Print("Please enter your weight(kg): ")
	userWeight, _ := reader.ReadString('\n')

	fmt.Print("Please enter your height(m): ")
	userHeight, _ := reader.ReadString('\n')

	fmt.Print(userWeight)
	fmt.Print(userHeight)
}
