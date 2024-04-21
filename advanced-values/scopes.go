package main

import "fmt"

var username = "Amir" // global scope because it is main package

func scope() {
	shouldContinue := true

	username := "Amirhosein" // this will shadow higher-level scoped variable

	if shouldContinue {
		userAge := 31

		fmt.Printf("Name: %v, Age: %v", username, userAge)
	}

	// fmt.Println(userAge) ---> doesn't work, cannot access block-scope outside it
}
