package main

import (
	"bufio"
	"fmt"
	"os"
)

type Storer interface {
	Store(fileName string)
}

type Prompter interface {
	PromptForInput()
}

type PromptStorer interface {
	Storer   // this is embedding
	Prompter // we can embed as many interfaces as we need to
}

type userInputData struct {
	input string
}

type user struct {
	firstName      string
	lastName       string
	*userInputData // ---> we can embed structs too, but we embed the pointer of the struct
}

func newUser(first string, last string) *user {
	return &user{first, last, &userInputData{}}
}

func newUserInputData() *userInputData {
	return &userInputData{} // initializing input string with empty value
}

func (usr *userInputData) PromptForInput() {
	fmt.Print("Your input please: ")

	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("fetching user input failed")
		return
	}

	usr.input = userInput
}

func (usr *userInputData) Store(fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("creating file failed")
		return
	}

	defer file.Close()
	file.WriteString(usr.input)
}

func main() {
	data := newUserInputData()
	// data.PromptForInput()
	// data.Store("user.txt")

	handleUserInput(data)

	user1 := newUser("Amirhosein", "Lashkari")
	user1.PromptForInput()
	user1.Store("user.txt")
}

func handleUserInput(container PromptStorer) {
	fmt.Println("Ready to store data")
	container.PromptForInput()
	container.Store("user.txt")
}
