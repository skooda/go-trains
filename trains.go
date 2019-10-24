package main

import (
	"fmt"
	"strconv"
	"sync"
)

type train struct {
	id int
}

func platform(s string, wg *sync.WaitGroup, mq chan train) {
	train := <-mq
	fmt.Println("Train " + strconv.Itoa(train.id) + " arrived to platform " + s)
	wg.Done()
}

func main() {
	// Init message queue
	messageQueue := make(chan train)

	// Add waitgroup as counter for coroutines
	var wg sync.WaitGroup
	defer wg.Wait()

	// Add some coroutines
	go platform("one", &wg, messageQueue)
	go platform("two", &wg, messageQueue)
	go platform("three", &wg, messageQueue)
	wg.Add(3) // ... and add them into counter

	// Send some trains into queue
	messageQueue <- train{id: 1}
	messageQueue <- train{id: 2}
	messageQueue <- train{id: 3}

}
