package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type logData struct {
	message  string
	fileName string
}

func (lData *logData) log() {
	err := ioutil.WriteFile(lData.fileName, []byte(lData.message), fs.FileMode(0644))

	if err != nil {
		fmt.Println("Storing the log data failed")
		message := err.Error() // Error is defined inside WriteFile interface
		fmt.Println(message)   // real Error message
	}
}

type loggableString string

// interface is like a contract
// VERY IMPORTANT: Go interfaces are only about methods
// naming convention: if interface has only one method its name should be: {method name} + "er"
type logger interface {
	log()
	// printMessage(message string) int ---> we can pass method requirements
	// message string ---> couldn't do this because it is not a method
}

func (text loggableString) log() {
	fmt.Println(text)
}

func interfaces() {
	userLog := logData{"user logged in", "user-log.txt"}
	// userLog.log()
	// createLog(userLog) ---> although it has log() method but it won't work because not type "loggableString"
	Logger(&userLog) // ---> will work because it has log() method

	message := loggableString("[INFO] user created")
	// message.log()
	createLog(message) // ---> will work because it is type "loggableString"
	Logger(message)    // ---> will work because it has log() method

	createLog("TEXT") // ---> again will work because it is type string that "loggableString" is (here is the problem)
	// Logger("TEXT") ---> will not work

	// fmt.Println(double(2))
	// fmt.Println(double(2.87))
	// fmt.Println(double(2.8734))
	// fmt.Println(double([]int{2, 3, 4, 5, 6}))
}

func createLog(data loggableString) {
	data.log()
}

func Logger(data logger) {
	data.log()
}

// we have empty interfaces for when we don't care about the type of input value
// IMPORTANT: empty interface means "any" type
func outputValue(value interface{}) {
	fmt.Println(value)

	val1, ok1 := value.(loggableString) // checks if a value is in a certain type
	val2, ok2 := value.(string)
	fmt.Println(val1, ok1, val2, ok2)
}

// writing a generic function with dynamic execution
func double(value any) interface{} {
	switch val := value.(type) {
	case int:
		return val * 2
	case float64:
		return val * 2
	case []int:
		newNumbers := append(val, val...)
		return newNumbers
	default:
		return ""
	}
}
