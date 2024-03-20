package main

import "fmt"

type person struct {
	name string
	age  int
}

type customNumber int // IMPORTANT: we defined custom type to be able to define method for it

// VERY IMPORTANT: defining a method on a type(customNumber) below:
func (num customNumber) power(power int) customNumber {
	result := num

	for i := 1; i < power; i++ {
		result = result * num
	}

	return result
}

type peopleType map[string]person

func customtype() {
	var people1 map[string]person = map[string]person{
		"1": {"Amir", 31},
		"2": {"Amirhosein", 30},
		"3": {"Hosein", 29},
	}

	var people2 peopleType = peopleType{
		"1": {"Amir", 31},
		"2": {"Amirhosein", 30},
		"3": {"Hosein", 29},
	}
	fmt.Println(people1, people2)
	// ---------------------------------------
	var newNum customNumber = 5
	poweredNum := newNum.power(3)
	fmt.Println(poweredNum)
}
