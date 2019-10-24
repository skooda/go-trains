package main

import (
	"fmt"
	"sync"
)

func platform(s string, wg *sync.WaitGroup) {
	fmt.Println("Train arrived to platform " + s)
	wg.Done()
}

func main() {

	// Add waitgroup as counter for coroutines
	var wg sync.WaitGroup
	defer wg.Wait()

	// Add some coroutines
	go platform("one", &wg)
	go platform("two", &wg)
	go platform("three", &wg)
	wg.Add(3) // ... and add them into counter

}
