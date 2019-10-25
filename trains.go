package main

import (
	"fmt"
	"strconv"
	"time"
)

type train struct {
	id int
}

type station struct {
	name string
}

func platform(station station, id int, arrival chan train, departure chan train) {
	for true { // Persistent platform
		train := <-arrival
		fmt.Println(station.name + ": Train " + strconv.Itoa(train.id) + " arrived to platform " + strconv.Itoa(id))
		time.Sleep(time.Second) // train waiting in the station
		fmt.Println(station.name + ": Train " + strconv.Itoa(train.id) + " leaving platform " + strconv.Itoa(id))
		departure <- train
	}
}

// track represents an rail between two stations
func track(arrival chan train, departure chan train) {
	train := <-arrival

	fmt.Println("Train " + strconv.Itoa(train.id) + " is on the way")
	time.Sleep(5 * time.Second) // simulate train journey
	fmt.Println("Train " + strconv.Itoa(train.id) + " is near the station")


	departure <- train // train leaves the track and arrives some station
}

func main() {
	maxTrainsOnTrack := 10

	// Add some stations
	prague := station{name: "Prague"}
	paris := station{name: "Paris"}

	// Init message queue
	pragueIn := make(chan train, maxTrainsOnTrack - 1)
	pragueOut := make(chan train, maxTrainsOnTrack - 1)
	parisIn := make(chan train, maxTrainsOnTrack - 1)
	parisOut := make(chan train, maxTrainsOnTrack - 1)

	// Add some platforms as coroutines
	go platform(prague, 1, pragueIn, pragueOut)
	go platform(prague, 2, pragueIn, pragueOut)
	go platform(paris, 1, parisIn, parisOut)
	go platform(paris, 2, parisIn, parisOut)
	go platform(paris, 3, parisIn, parisOut)

	go track(pragueOut, parisIn)
	go track(parisOut, pragueIn)

	// Send some trains into queue
	trainCount := 5

	for i := 0; i < trainCount; i+=2 {
		pragueIn <- train{id: i} // Some trains is starting in prague
		parisIn <- train{id: i+1} // .. and some in paris
	}

	time.Sleep(60 * time.Second) // Disabled wg and just run the simulation for one minute for now
}
