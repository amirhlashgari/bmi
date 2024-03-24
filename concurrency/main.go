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
	limiter := make(chan int, 3) // three concurrent tasks in a while(to avoid over performing)
	// NOTE: channels would be free after sending a value out of a channel

	// x := go generateValue(c) ---> not allowed so we should use channels to return our value
	go generateValue(c, limiter)
	go generateValue(c, limiter)

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

func generateValue(c chan int, limit chan int) int {
	limit <- 1 // ---> after three times writing 1 to channels it will stop till one ends
	sleepTime := randN.Intn(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	result := rand.Intn(10)
	c <- result
	<-limit       // setting channel free
	return result // not neccessary
}
