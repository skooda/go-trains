package main

import (
	"fmt"
	"strconv"
	"sync"
)

type train struct {
	id int
}

type station struct {
	name string
}

func platform(station station, id int, wg *sync.WaitGroup, mq chan train) {
	for true {
		train := <-mq
		fmt.Println(station.name + ": Train " + strconv.Itoa(train.id) + " arrived to platform " + strconv.Itoa(id))
		wg.Done()
	}
}

func main() {
	// Init message queue
	messageQueue := make(chan train)

	// Add waitgroup as counter for coroutines
	var wg sync.WaitGroup
	defer wg.Wait()

	// Add some stations
	prague := station{name: "Prague"}
	paris := station{name: "Paris"}

	// Add some platforms as coroutines
	go platform(prague, 1, &wg, messageQueue)
	go platform(prague, 2, &wg, messageQueue)
	go platform(paris, 1, &wg, messageQueue)
	go platform(paris, 2, &wg, messageQueue)
	go platform(paris, 3, &wg, messageQueue)

	// Send some trains into queue
	trainCount := 10
	wg.Add(trainCount) // ... and add them into counter

	for i := 0; i < trainCount; i++ {
		messageQueue <- train{id: i}
	}

}
