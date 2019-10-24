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

	var wg sync.WaitGroup

	go platform("one", &wg)
	go platform("two", &wg)
	go platform("three", &wg)
	wg.Add(3)

	wg.Wait()
}
