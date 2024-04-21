package main

import (
	"bufio"
	"fmt"
	"os"
)

func deferpanic() {
	userInput := getUserInput()
	storeData(userInput)
}

func getUserInput() string {
	fmt.Println("Enter the text should be stored")
	fmt.Print("Input: ")

	reader := bufio.NewReader(os.Stdin)
	enteredText, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Failed to read user input")
		return ""
	}

	return enteredText
}

func storeData(data string) {
	file, err := os.Create("data.txt")

	if err != nil {
		fmt.Println("Creating the file failed")
		panic(err) // because panic() crashes go program execution you should use it with caution
	}

	defer file.Close() // ---> this way we can't have error handling. to do so we should do as below:
	/**
	 * defer func() {
			err := file.Close()
			if err != nil {
				fmt.Println("Closing the file failed")
			}
	 	}()
	 *
	 * NOTE: look at anonymous functiuon execution right after defining it func(){}()
	*/

	file.WriteString(data)
	fmt.Println("Successfully stored data in file")
	/**
	 * instead of using file.close() here to avoid forgetting this we put it on top of function
	 * then using "defer" keyword we tell function to execute this line at the end
	 */
	// file.Close()
}
