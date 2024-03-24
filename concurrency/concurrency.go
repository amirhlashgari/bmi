package main

import (
	"fmt"
	"os"
)

func concurrency() {
	greet()
	storeData("dummy data", "data.txt")

	channel := make(chan int)

	go storeMoreData(50000, "50000_1.txt", channel) // this is how we use goroutines
	go storeMoreData(50000, "50000_2.txt", channel)

	<-channel // channels send signals to go program that when would go routine is done
	<-channel // second one is for waiting to second function be done
}

func greet() {
	fmt.Println("Hello world")
}

func storeData(storableText string, fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("Creating the file failed")
		return
	}

	defer file.Close()

	_, err = file.WriteString(storableText)
	if err != nil {
		fmt.Println("writing to the file failed")
	}
}

func storeMoreData(lines int, fileName string, c chan int) {
	for i := 0; i < lines; i++ {
		text := fmt.Sprintf("Line %v - Dummy Data\n", i)
		storeData(text, fileName)
	}

	fmt.Printf("-- Done storing %v lines of text --\n", lines)
	c <- 1
}
