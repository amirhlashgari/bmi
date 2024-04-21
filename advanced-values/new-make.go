package main

import "fmt"

func makefunc() {
	sliceHobbies := []string{"Sport", "Reading"}               // creates a slice
	sliceHobbies = append(sliceHobbies, "Cooking", "Watching") // this will make new array and assigns it again to sliceHobbies

	/**
	* the problem with creating slice like above is that when we want to append, it will make new array and add new items to it
	*
	* by using "make" we can avoid creating new array since we will tell slice capacity in creation phase
	* NOTE: hobbies is not limited to be only up to 10 items. if we want to append more than 10 items to it after that append will make new array
	*
	* "make" is not just for Slices, we can use it in Maps and also Channels
	* mapExample := make(map[string]int, 5)
	 */

	hobbies := make([]string, 2, 10)
	hobbies[0] = "sports"
	hobbies[1] = "reading"
	// hobbies[2] = "cooking" ---> this will return error because it is initiated with 2 values, unless we append it
	hobbies = append(hobbies, "cooking", "watching")

	fmt.Println(hobbies)
}

func newfunc() {
	// "new"(rarely used) allocates memory for the value without initializing it and returns pointer
	hobbies := new([]string)
	*hobbies = append(*hobbies, "test", "test2", "test3", "test4")
	fmt.Println(*hobbies)

	number := new(int)   // it is equal to ( number := 0 )
	fmt.Println(number)  // returns pointer's address  ----> ( fmt.Println(&number) )
	fmt.Println(*number) // returns zero (0) - in case we want to create a zero
}
