package main

import (
	"fmt"
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().Unix())
var randN = rand.New(source)

func main() {
	c := make(chan int)

	// x := go generateValue(c) ---> not allowed so we should use channels to return our value
	go generateValue(c)
	go generateValue(c)

	x := <-c
	y := <-c

	sum := x + y
	/**
	 *  IMPORTANT: another way to do above(looping through channels):

	 	go generateValue(c)
		go generateValue(c)
		go generateValue(c)
		go generateValue(c)

		sum := 0
		iteration := 4      ---> number of function calls

		for num := range c {
			sum += num
			i++
			if i == 4 {
				close(c)	---> when looping to channels we should be aware to close them
			}
		}
	*/

	fmt.Println(sum)
}

func generateValue(c chan int) int {
	sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	result := rand.Intn(10)
	c <- result
	return result // not neccessary
}
