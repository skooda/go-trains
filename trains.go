package main

import (
	"fmt"
	"sync"
)

func platform(s string, wg *sync.WaitGroup, mq chan string) {
	fmt.Println(<-mq + " arrived to platform " + s)
	wg.Done()
}

func main() {
	// Init message queue
	messageQueue := make(chan string)

	// Add waitgroup as counter for coroutines
	var wg sync.WaitGroup
	defer wg.Wait()

	// Add some coroutines
	go platform("one", &wg, messageQueue)
	go platform("two", &wg, messageQueue)
	go platform("three", &wg, messageQueue)
	wg.Add(3) // ... and add them into counter

	// Send some trains into queue
	messageQueue <- "Train 1"
	messageQueue <- "Train 2"
	messageQueue <- "Train 3"

}
