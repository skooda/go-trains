package main

import (
	"fmt"
	"time"
)

func platform(s string) {
	fmt.Println("Train arrived to platform " + s)
}

func main() {
	go platform("one")
	go platform("two")
	go platform("three")
	time.Sleep(time.Second)

}
