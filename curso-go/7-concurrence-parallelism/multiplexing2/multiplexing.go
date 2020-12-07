package main

import (
	"fmt"
	"time"
)

//Methdo receives a string and returns read-only string channel
// Has a internal generator
func speak(person string) <-chan string {
	c := make(chan string)
	go func() { //Anonimous function with concurrency execution
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s speaking: %d\n", person, i+1)
		}
	}()

	return c
}

// Multiplexing/joining info of two channels and returning it in a single read-only channel
func multiplex(entry1, entry2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-entry1:
				c <- s
			case s := <-entry2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := multiplex(speak("John"), speak("Mary"))
	for i := 0; i < 6; i++ {
		fmt.Println(<-c)
	}
}
