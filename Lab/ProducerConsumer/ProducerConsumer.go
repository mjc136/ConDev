package main

import (
	"sync"
)

const (
	numThreads = 100
	size       = 20
)

// consumer gets events from the channel and processes them
func Consumer() {

}

// producer creates events and sends them to the channel
func Producer() {

}

func main() {
	var wg sync.WaitGroup
	theChan := make(chan bool) //use unbuffered channel in place of semaphore

	wg.Wait()
	close(theChan)
}
