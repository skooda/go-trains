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
	train := <-mq
	fmt.Println(station.name + ": Train " + strconv.Itoa(train.id) + " arrived to platform " + strconv.Itoa(id))
	wg.Done()
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
	wg.Add(5) // ... and add them into counter

	// Send some trains into queue
	messageQueue <- train{id: 1}
	messageQueue <- train{id: 2}
	messageQueue <- train{id: 3}
	messageQueue <- train{id: 4}
	messageQueue <- train{id: 5}

}
